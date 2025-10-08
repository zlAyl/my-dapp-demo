package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/zlAyl/my-dapp-demo/task1/count"
)

func main() {
	//1.获取env文件
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envPath := filepath.Join(dir, "task1", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
	}
	rpcUrl := os.Getenv("RPC_URL")

	//创建client
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error converting public key to ECDSA")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//初始化交易opt实例
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)     // in wei
	//auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = gasPrice

	//部署合约
	//address, _, _, err := count.DeployCount(auth, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("合于地址： %s\n", address.Hex())

	address := common.HexToAddress("0xfdd72bee132897EBF8f29774E644d07DBbB1a0D2")

	//fmt.Printf("Deployed contract： %s\n", tx.Hash().Hex())

	//加载合约
	countContract, err := count.NewCount(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//调用合约方法
	tx, err := countContract.AddCount(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	// 等待交易被确认
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatalf("交易失败: %s", tx.Hash().Hex())
	}

	callOpt := &bind.CallOpts{Context: context.Background()}
	nextCount, err := countContract.Count(callOpt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("next count:", nextCount.String())

}
