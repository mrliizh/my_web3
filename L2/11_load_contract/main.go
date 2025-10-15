package main

import (
	"fmt"
	"log"

	"example.com/init_project/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddr = "0x0Ba442FFa124333c9F82f680A3230e46CbeB1DAA"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/5PoteT7uHuIrP414jmfKWNcvDwwXrpmA")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	_ = storeContract

    fmt.Println(storeContract)

}