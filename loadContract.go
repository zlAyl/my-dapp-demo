package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/zlAyl/my-dapp-demo/store"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//加载合约
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := "0xCaA2BC69381DC26174F88fFfaa89C189802319b1"

	newStore, err := store.NewStore(common.HexToAddress(contractAddress), client)
	if err != nil {
		return
	}

	_ = newStore
}
