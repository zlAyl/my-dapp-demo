package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envPath := filepath.Join(dir, "task1", ".env")
	fmt.Println(envPath)
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
	}
	rpcUrl := os.Getenv("RPC_URL")
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//获取最新区块信息
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return
	}

	fmt.Println("区块hash:", block.Hash().Hex())
	fmt.Println("区块高度：", block.Number())
	fmt.Println("区块时间:", block.Time())
	fmt.Println("Nonce:", block.Nonce())
	fmt.Println("GasLimit:", block.GasLimit())
	fmt.Println("header time:", block.Header().Time)
	fmt.Println("交易数量:", len(block.Transactions()))
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		return
	}
	fmt.Println("交易数量:", count)
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return
	}
	fmt.Println("ChainID:", chainId)

	//发送ETH转账交易

	//1.先获取fromAddress
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress:", fromAddress.Hex())

	//获取nonce
	//发送新交易时使用包含了没有打包的交易nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}
	//client.NonceAt(context.Background(), fromAddress, nil)   //查询历史或者指定区块高度的nonce
	value := big.NewInt(1000000000000000000) //1ETH

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xF03399DB09412a3298b671be774f849eB747163e")
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		Value:    value,
		GasPrice: gasPrice,
		Gas:      uint64(21000),
		To:       &toAddress,
		Data:     nil,
	})
	//交易签名
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		return
	}
	if err := client.SendTransaction(context.Background(), signTx); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signTx.Hash().Hex())

}
