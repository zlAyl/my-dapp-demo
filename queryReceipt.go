package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	blochNumber := big.NewInt(9372875)
	blockHash := common.HexToHash("0x1449083045df72a7744cab27a03859ed9d078ffa71c1cf5ab8d6983f5a1dfa8a")
	//获取区块下得收据列表 根据blockHash
	receipts, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	//获取区块下得收据列表 根据blockNumber
	blockReceipts, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blochNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipts[0])
	fmt.Println(blockReceipts[0])

	for _, receipt := range receipts {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.TxHash.Hex())
		fmt.Println(receipt.TransactionIndex)
		fmt.Println(receipt.ContractAddress.Hex())

		break
	}

	txHash := common.HexToHash("0xde7084aec5e33db03e2141c7553d7f749d1226e741b2588a07b38b2a5079a4b1")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.TxHash.Hex())
}
