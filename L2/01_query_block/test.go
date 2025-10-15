package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
        client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/5PoteT7uHuIrP414jmfKWNcvDwwXrpmA")  //API Key用来验证身份和控制流量的,具体分配哪台节点，由 Alchemy 的负载均衡系统自动决定
        if err != nil {
                log.Fatal(err)
        }

        blockNumber := big.NewInt(987918)

        header, err := client.HeaderByNumber(context.Background(), blockNumber)
        fmt.Println(header.Number.Uint64())
        fmt.Println(header.Time)
        

        block, err := client.BlockByNumber(context.Background(), nil)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(block.Number().Uint64())     
        fmt.Println(block.Time())       
        timestamp := time.Unix(int64(block.Time()), 0)
        fmt.Println(timestamp)
        fmt.Println(block.Difficulty().Uint64()) 
        fmt.Println(block.Hash().Hex())           
        fmt.Println(len(block.Transactions()))    

        count, err := client.TransactionCount(context.Background(), block.Hash())
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(count) // 144
}