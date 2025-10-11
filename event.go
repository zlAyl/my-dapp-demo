package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {

}

func main() {
	storeABI := `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
	contractAbi, err := abi.JSON(strings.NewReader(storeABI))
	if err != nil {
		log.Fatal(err)
	}
	//查询事件
	{
		client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
		if err != nil {
			log.Fatal(err)
		}

		contractAddress := "0xCaA2BC69381DC26174F88fFfaa89C189802319b1"
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(9382228),
			//ToBlock:   big.NewInt(0),
			Addresses: []common.Address{
				common.HexToAddress(contractAddress),
			},
			//Topics: [][]common.Hash{},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			fmt.Println(vLog.BlockNumber)
			//通过Abi来解码数据

			_, exists := contractAbi.Events["ItemSet"]
			if !exists {
				fmt.Printf("事件 ItemSet  不存在")
			}

			//第一种方式
			result := make(map[string]interface{})
			err := contractAbi.UnpackIntoMap(result, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			if keyBytes, ok := result["key"].([32]byte); ok {
				//keyHex := common.Bytes2Hex(keyBytes[:])
				//keyStr := string(bytes.Trim(keyBytes[:], "\x00"))
				// 尝试转换为原始字符串（如果数据原本是字符串）
				keyStr := string(bytes.Trim(keyBytes[:], "\x00"))
				fmt.Printf("Key (string): %s\n", keyStr)
			}

			if valueBytes, ok := result["value"].([32]byte); ok {
				//keyHex := common.Bytes2Hex(keyBytes[:])
				//keyStr := string(bytes.Trim(keyBytes[:], "\x00"))
				// 尝试转换为原始字符串（如果数据原本是字符串）
				valStr := string(bytes.Trim(valueBytes[:], "\x00"))
				fmt.Printf("Value (string): %s\n", valStr)
			}

			//第2种方式
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err = contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Key:", string(bytes.Trim(event.Key[:], "\x00")))
			fmt.Println("Value:", string(bytes.Trim(event.Value[:], "\x00")))
		}
	}

	//订阅事件
	{
		eventClient, err := ethclient.Dial("wss://ethereum-sepolia-rpc.publicnode.com")
		if err != nil {
			log.Fatal(err)
		}
		contractAddress := common.HexToAddress("0xCaA2BC69381DC26174F88fFfaa89C189802319b1")
		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		}

		logs := make(chan types.Log)

		sub, err := eventClient.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case err := <-sub.Err():
				fmt.Println(err)
			case vlog := <-logs:
				fmt.Println(vlog.BlockNumber)
				event := struct {
					Key   [32]byte
					Value [32]byte
				}{}
				err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vlog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("sub key:", string(bytes.Trim(event.Key[:], "\x00")))
				fmt.Println("sub value:", string(bytes.Trim(event.Value[:], "\x00")))

				//读取topic  第一个主题总是事件的签名。
				for _, v := range vlog.Topics {
					fmt.Println(v.Hex())
				}
			}

		}
	}
}
