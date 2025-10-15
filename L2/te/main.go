package main

import (
	"context"
	"crypto/ecdsa"

	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main(){
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/5PoteT7uHuIrP414jmfKWNcvDwwXrpmA")  //API Key用来验证身份和控制流量的,具体分配哪台节点，由 Alchemy 的负载均衡系统自动决定
	if err != nil{
		log.Fatal(err)
	}
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	BlockReceipts, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil{
		log.Fatal(err)
	}
	for _, receipt := range BlockReceipts{
		fmt.Println(receipt.Status)
		break
	}

	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println((hexutil.Encode(privateKeyBytes[2:])))

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok{
		log.Fatal("")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes[4:]))

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(address)

	nonce, err := client.PendingNonceAt(context.Background(), address)
	fmt.Println(nonce)
}