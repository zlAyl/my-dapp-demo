package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	//获取env
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envPath := filepath.Join(dir, "task1", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
	}
	//初始化client
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	//合约地址(从部署合约后拿到)
	contractAddressHex := "0xfdd72bee132897EBF8f29774E644d07DBbB1a0D2"
	contractAddress := common.HexToAddress(contractAddressHex)

	//加载合约
	//countContract, err := count.NewCount(contractAddress, client)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//获取合约事件
	//初始化查询
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(9361823),
		Addresses: []common.Address{contractAddress},
	}

	//如果不确定fromBlock 可以根据创建合约时候的交易hash 来获取交易时的blockNumber
	//可以从网页找到创建的合约拿到交易哈希 0xda5f76e23c3bee8717f0f808d944ae24a0f31ebedd1e50e8981987fb443637aa

	//从部署合约时区块查找  9361823
	//deployTxHash := "0xda5f76e23c3bee8717f0f808d944ae24a0f31ebedd1e50e8981987fb443637aa"
	//
	//_, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(deployTxHash))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if !isPending {
	//	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(deployTxHash))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("创建合约时候的 blockNumber:", receipt.BlockNumber)
	//}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	countAbi := `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"nextCount","type":"uint256"}],"name":"AddCount","type":"event"},{"inputs":[],"name":"addCount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"count","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

	contractAbi, err := abi.JSON(strings.NewReader(countAbi))
	if err != nil {
		log.Fatal(err)
	}

	_, exists := contractAbi.Events["AddCount"]
	if !exists {
		fmt.Printf("事件 AddCount  不存在")
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		//第一种方式
		//result := make(map[string]interface{})
		//err := contractAbi.UnpackIntoMap(result, "AddCount", vLog.Data)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println(result)
		//
		//for key, value := range result {
		//	fmt.Println(key, value)
		//}

		//第2种方式
		event := struct {
			NextCount *big.Int
		}{}
		if err := contractAbi.UnpackIntoInterface(&event, "AddCount", vLog.Data); err != nil {
			log.Fatal(err)
		}
		fmt.Println(event.NextCount.String())

		// 处理 Topics（索引参数）
		for i, topic := range vLog.Topics {
			if i > 0 {
				fmt.Printf("Topic %d: %s\n", i, topic.Hex())
				//地址（address）类型在Solidity中占20字节（160位）。但是在日志的数据部分，每个参数都会填充到32字节。
				//所以一个地址在数据部分中会以32字节的形式存在，其中前12字节是0，后20字节是实际的地址。
				sendAddress := common.Bytes2Hex(topic[12:32])
				fmt.Printf("sendAddress %d: %s\n", i, sendAddress)
			}
		}
		// 处理 Data（非索引参数）
		//nextCount := common.Bytes2Hex(vLog.Data[:])
		//nextCount := new(big.Int).SetBytes(vLog.Data[0:32])
		//fmt.Printf("数据: %x\n", nextCount.String())
		//fmt.Println("---")
	}
	//if len(logs) > 0 {
	//	// 第一个日志的交易哈希可能就是部署交易
	//	fmt.Printf("交易hash %v\n", logs[0].TxHash)
	//}
}
