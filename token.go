package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//代币转账
	//A 钱包 发送交易到 代币合约
	//
	//代币合约 收到交易后，执行 transfer 函数
	//
	//代币合约 内部更新余额：
	//
	//A 的余额减少
	//
	//B 的余额增加
	//
	//交易完成，代币从 A 转移到 B

	//先部署代币合约 MyToken
	//部署后的地址为 0xefEaf7Bf78C0c13473359504efa7e4AF3951B598

	//client
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	//获取私钥 然后从私钥推算出公钥  再推算出formAddress
	pk := os.Getenv("PRIVATE_KEY")

	privateKeyECDSA, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKeyECDSA.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress("0xF03399DB09412a3298b671be774f849eB747163e")    //接收代币转账的地址
	tokenAddress := common.HexToAddress("0xefEaf7Bf78C0c13473359504efa7e4AF3951B598") //代币合约地址

	//构建交易数据
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//交易data 调用合约的 transfer(ERC20)
	//再不使用合于abi的情况下手动构建data
	//ERC20的transfer函数签名为：transfer(address to, uint256 value)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4] //前4位
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	amount := big.NewInt(1000000000000000000)
	//amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...) //添加接收方地址参数 (32字节)
	data = append(data, paddedAmount...)  // 添加金额参数 (32字节)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress, //代币合约地址
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		Data:     data,
		Gas:      gasLimit,
	})

	//交易签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		return
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatal(receipt.Status)
	}
	fmt.Printf("tx hash: %s\n", signedTx.Hash().Hex())

	//调用合约balanceOf 查询余额 因为查询余额不需要消耗gas 所以不通过交易来获取 所以不能使用SendTransaction
	//CallContract 是用来查询的 不改变状态
	//transferFnSignature := []byte("balanceOf(address)")
	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(transferFnSignature)
	//methodID := hash.Sum(nil)[:4] //前4位
	//paddedAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
	//var data []byte
	//data = append(data, methodID...)
	//data = append(data, paddedAddress...) //添加接收方地址参数 (32字节)
	//
	//result, err := client.CallContract(context.Background(), ethereum.CallMsg{
	//	To:   &tokenAddress,
	//	Data: data,
	//}, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("转账账户余额：%v\n", new(big.Int).SetBytes(result))

}
