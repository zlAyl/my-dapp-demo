package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//执行合约

	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := "0xCaA2BC69381DC26174F88fFfaa89C189802319b1"

	//1.使用生成的go代码
	//{
	//newStore, err := store.NewStore(common.HexToAddress(contractAddress), client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var key [32]byte
	//var value [32]byte
	////执行合约
	//copy(key[:], []byte("demo_save_key"))
	//copy(value[:], []byte("demo_save_value22222"))
	//
	//privateKey := os.Getenv("PRIVATE_KEY")
	//privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//chainID, err := client.ChainID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//txOpts, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, chainID)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//tx, err := newStore.SetItem(txOpts, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("setItem:", tx.Hash().Hex())
	//
	//callOpts := &bind.CallOpts{Context: context.Background()}
	//item, err := newStore.Items(callOpts, key)
	//if err != nil {
	//	return
	//}
	//fmt.Println("items:", string(item[:]))
	//}

	//2.使用 abi 文件调用合约
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

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	storeAbi := `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
	// 准备交易数据
	abiJson, err := abi.JSON(strings.NewReader(storeAbi))
	if err != nil {
		log.Fatal(err)
	}
	methodName := "setItem"
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key_use_abi"))
	copy(value[:], []byte("demo_save_value_use_abi_11111")) //不够32位会后面自动补0
	input, err := abiJson.Pack(methodName, key, value)
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(contractAddress)
	tx := types.NewTx(&types.LegacyTx{
		To:       &toAddress,
		Nonce:    nonce,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		Gas:      300000,
		Data:     input,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal(err)
	}

	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatal("failed to set item")
	}

	fmt.Println("交易hash:", signedTx.Hash().Hex())

	//查询数据
	callInput, err := abiJson.Pack("items", key)
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Data: callInput,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var unpacked [32]byte
	if err := abiJson.UnpackIntoInterface(&unpacked, "items", result); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(unpacked[:]))
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
}
