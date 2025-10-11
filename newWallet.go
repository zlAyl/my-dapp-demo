package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	//生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("privateKey:", privateKey)
	//0xc92D410CC2C94757DF417790FEB1D62c5a8F783b
	//620ba2596941ce70de066c5ad836531dbff27880b8429f53f1db3e8c9c7aeec7
	//b3daecbacc9ee093df02e5924843a2b7c9027aae81f2ce5b2a0ebbad456522e9

	privateKeyBytes := crypto.FromECDSA(privateKey)
	pk := hexutil.Encode(privateKeyBytes)[2:] // 去掉'0x' 就是钱包的私钥
	fmt.Println("privateKey:", pk)

	//根据私钥货得公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	//直接调用内置方法生成公钥
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress:", address)

	//还可以手动获取公钥
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println("这是公钥数据 不是地址:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:]) // 去掉'0x04'
	hashBytes := hash.Sum(nil)     // 32 字节
	addressBytes := hashBytes[12:] // 20 字节
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	//公共地址其实就是公钥的 Keccak-256 哈希，然后我们取最后 40 个字符（20 个字节）并用“0x”作为前缀
	fmt.Println(hexutil.Encode(addressBytes)) //bytes 转换成 16进制 string 实际以太坊地址

	//// 1. 获取原始公钥 (65字节，以0x04开头)
	//publicKeyBytes := crypto.FromECDSAPub(&privateKey.PublicKey)
	//// 例如: 0x04a1b2c3d4... (65字节)
	//
	//// 2. 去掉0x04前缀，得到64字节的X+Y坐标
	//publicKeyData := publicKeyBytes[1:]  // 64字节
	//
	//// 3. 计算Keccak-256哈希
	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(publicKeyData)  // 对64字节数据进行哈希
	//hashResult := hash.Sum(nil)  // 得到32字节哈希
	//
	//// 4. 取最后20字节作为地址
	//addressBytes := hashResult[12:]  // 20字节
	//ethereumAddress := common.BytesToAddress(addressBytes)
}
