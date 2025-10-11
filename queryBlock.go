package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//client 客户端
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	//区块头 参数为nil 时候 是获取最新得区块头
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return
	}
	fmt.Printf("最新区头hash: %v\n", header.Hash().Hex())
	fmt.Printf("最新区块高度: %v\n", header.Number.Uint64())
	fmt.Printf("最新区块Nonce: %v\n", header.Nonce.Uint64()) //pos 此值为0
	fmt.Printf("最新区块 gasLimit: %v\n", header.GasLimit)
	fmt.Printf("最新区块 Difficulty: %v\n", header.Difficulty.Uint64()) ////pos 此值为0
	fmt.Printf("最新区块 time: %v\n", time.Unix(int64(header.Time), 0).Format("2006-01-02 15:04:05"))

	blockNumber := big.NewInt(9372875)
	//获取block
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Block hash: %v\n", block.Hash().Hex())
	fmt.Printf("Block header hash: %v\n", block.Header().Hash().Hex())
	fmt.Printf("Block number: %v\n", block.Number().Uint64())
	fmt.Printf("Block time: %v\n", time.Unix(int64(block.Time()), 0))
	fmt.Printf("Block difficulty: %v\n", block.Difficulty().Uint64())
	fmt.Printf("Block len txs: %v\n", block.Transactions().Len())

}
