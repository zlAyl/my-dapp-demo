package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	privateKey := os.Getenv("PRIVATE_KEY")

	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKeyEcdsa.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fomAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//查询账户余额 返回指定区块高度（或区块哈希）的账户余额。这个余额是已经确认的，
	balance, err := client.BalanceAt(context.Background(), fomAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	log.Printf("Balance: %v", ethValue)

	//返回当前待处理状态下的账户余额 这个余额考虑了内存池中尚未被纳入区块的交易，即包括待确认的交易对余额的影响。
	at, err := client.PendingBalanceAt(context.Background(), fomAddress)
	if err != nil {
		return
	}
	fAt := new(big.Float)
	fAt.SetString(at.String())
	ethValue = new(big.Float).Quo(fAt, big.NewFloat(math.Pow10(18)))
	log.Printf("PendingBalance: %v", ethValue)

}
