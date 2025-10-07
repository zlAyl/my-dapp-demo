package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/653c5ddbe0094c66af5291d4068f3679")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	//blockNumber := big.NewInt(5671744)
	//
	//header, err := client.HeaderByNumber(context.Background(), blockNumber) //传nil表示获取最新的区块
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("header: %v\n", header)
	//fmt.Printf("header number: %v\n", header.Number)
	//fmt.Printf("header time: %v\n", header.Time)
	//fmt.Printf("header gas: %v\n", header.GasLimit)
	//fmt.Printf("header gas: %v\n", header.GasUsed)
	//fmt.Printf("header nonce: %v\n", header.Nonce)
	//fmt.Printf("header difficulty: %v\n", header.Difficulty)
	//fmt.Printf("header hash: %v\n", header.Hash().Hex())
	//
	//block, err := client.BlockByNumber(context.Background(), blockNumber)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("block: %v\n", block)
	//fmt.Printf("block number: %v\n", block.Number())
	//fmt.Printf("block time: %v\n", block.Time())
	//fmt.Printf("block gas: %v\n", block.GasLimit())
	//fmt.Printf("block gas: %v\n", block.GasUsed())
	//fmt.Printf("block nonce: %v\n", block.Nonce())
	//fmt.Printf("block difficulty: %v\n", block.Difficulty())
	//fmt.Printf("block hash: %v\n", block.Hash().Hex())
	//fmt.Printf("block len Transactions: %v\n", len(block.Transactions()))
	//
	//count, err := client.TransactionCount(context.Background(), block.Hash())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("transaction count: %v\n", count)
	//
	//chainID, err := client.ChainID(context.Background()) //可以获取当前连接网络的ChainID
	//if err != nil {
	//	panic(err)
	//}
	////client.NetworkID(context.Background())
	//for i := uint(0); i < count; i++ {
	//	tx := block.Transactions()[i]
	//	//tx, err := client.TransactionInBlock(context.Background(), block.Hash(), i)
	//	//if err != nil {
	//	//	return
	//	//}
	//
	//	//client.TransactionByHash
	//	signer := types.NewEIP155Signer(chainID)
	//	sender, err := types.Sender(signer, tx)
	//	if err != nil {
	//		panic(err)
	//	}
	//	//fmt.Printf("sender: %s\n", sender)
	//	fmt.Printf("sender hex: %s\n", sender.Hex())
	//	//fmt.Printf("tx: %v\n", tx.Type())
	//	//types.DynamicFeeTxType //EIP -1559交易
	//	//types.AccessListTxType //EIP -2390交易
	//
	//	//获取交易收据
	//	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("receipt status:%v\n", receipt.Status) // 1 成功
	//	fmt.Printf("receipt logs:%v\n", receipt.Logs)     // ...
	//	break
	//}
	//
	//receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(block.Hash(), false))
	//if err != nil {
	//	panic(err)
	//}
	////receiptByNumber, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(block.Number().Uint64())))
	//
	//for _, receipte := range receiptByHash {
	//	fmt.Printf("receipt status:%v\n", receipte.Status)
	//	fmt.Printf("receipt hash:%v\n", receipte.TxHash.Hex())
	//	fmt.Printf("receipt hash:%v\n", receipte.TransactionIndex)
	//}
	//
	////创建钱包
	//privateKey, err := crypto.GenerateKey()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) //私钥 也就是密码  去掉'0x'
	//
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("不是 ECDSA 公钥")
	//}
	//
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	//address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(address)

	//ETH转账
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(1000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)             // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0xc92D410CC2C94757DF417790FEB1D62c5a8F783b")

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}
