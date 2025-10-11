package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	//发送ETH

	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("fromAddress: %s\n", fromAddress.Hex())

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(1000000000000000000) //in wei (1 eth)
	gasLimit := uint64(21000)                // in units

	toAddress := common.HexToAddress("0xc92D410CC2C94757DF417790FEB1D62c5a8F783b")

	//生成交易数据
	//types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil) 已弃用
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
	})

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//对交易数据签名
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//发送交易
	if err := client.SendTransaction(context.Background(), signTx); err != nil {
		log.Fatal(err)
	}

	//等待交易完成
	receipt, err := bind.WaitMined(context.Background(), client, signTx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatalf("交易失败: %s", tx.Hash().Hex())
	}

	fmt.Printf("tx sent: %s", signTx.Hash().Hex())

	// 0xefEaf7Bf78C0c13473359504efa7e4AF3951B598 合约地址
	//0xefeaf7bf78c0c13473359504efa7e4af3951b598

}
