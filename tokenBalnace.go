package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/zlAyl/my-dapp-demo/myToken"
	"golang.org/x/crypto/sha3"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress("0xF03399DB09412a3298b671be774f849eB747163e")
	tokenAddress := common.HexToAddress("0xefEaf7Bf78C0c13473359504efa7e4AF3951B598")

	//查询余额不使用abi文件
	//调用合约balanceOf
	transferFnSignature := []byte("balanceOf(address)")

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)

	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)

	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %v\n", new(big.Int).SetBytes(result))

	//使用abi文件 生成go文件
	//solcjs --abi  --include-path node_modules --base-path . contracts/MyToken.sol
	//abigen --abi=contracts_MyToken_sol_MyToken.abi  --pkg=myToken --out=myToken.go

	//合约实例
	instance, err := myToken.NewMyToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	//查询用户代币余额
	balance, err := instance.BalanceOf(&bind.CallOpts{}, fromAddress)
	if err != nil {
		return
	}

	fmt.Printf("balance: %v\n", balance.String())

}
