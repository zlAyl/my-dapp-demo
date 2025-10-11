package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//获取client
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(9372875)

	//获取链ID
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("链 ID: %v\n", chainId)

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block hash:", block.Hash().Hex())
	//获取交易
	//第一种  根据 block.Transactions 来获取
	for _, tx := range block.Transactions() {
		fmt.Println("交易 hash:", tx.Hash().Hex())
		fmt.Println("交易 gasPrice:", tx.GasPrice().Uint64())
		fmt.Println("交易 gas:", tx.Gas())
		fmt.Println("交易 time:", tx.Time().Format("2006-01-02 15:04:05"))
		fmt.Println("交易 data:", tx.Data())
		fmt.Println("交易 chainId:", tx.ChainId().Uint64())
		fmt.Println("交易 value:", tx.Value().Uint64())
		fmt.Println("交易 接收者:", tx.To().Hex())

		//根据交易获取交易得发送者
		sender, err := types.Sender(types.LatestSignerForChainID(chainId), tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("交易 sender:", sender.Hex())

		//交易收据
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("收据 status:", receipt.Status)
		fmt.Println("交易 logs:", receipt.Logs)
		break

	}

	//第二种  如果已知 交易得hash
	txHash := common.HexToHash("0xde7084aec5e33db03e2141c7553d7f749d1226e741b2588a07b38b2a5079a4b1")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	if !isPending {
		fmt.Println("交易hasH:", tx.Hash().Hex())
		fmt.Println("交易 gasPrice:", tx.GasPrice().Uint64())
		fmt.Println("交易 gas:", tx.Gas())
		fmt.Println("交易 time:", tx.Time().Format("2006-01-02 15:04:05"))
		fmt.Println("交易 data:", tx.Data())
		fmt.Println("交易 chainId:", tx.ChainId().Uint64())
		fmt.Println("交易 value:", tx.Value().Uint64())
		fmt.Println("交易 接收者:", tx.To().Hex())

		sender, err := types.Sender(types.LatestSignerForChainID(chainId), tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("交易 发送者:", sender.Hex())

	}

	//第三中 已经区块hash
	//blockHash := common.HexToHash("")
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	for i := uint(0); i < count; i++ {
		tx, err := client.TransactionInBlock(context.Background(), block.Hash(), i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
		break
	}
}
