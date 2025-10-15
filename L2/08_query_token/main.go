package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/init_project/erc20" // for demo
)
func main() {
    	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/5PoteT7uHuIrP414jmfKWNcvDwwXrpmA")
        if err != nil {
                log.Fatal(err)
        }
        // Golem (GNT) Address
        tokenAddress := common.HexToAddress("0x99a2F0456313D93E9E0260C7cC70f834Cf188Ed4")
        instance, err := token.NewErc20(tokenAddress, client)  //
        if err != nil {
                log.Fatal(err)
        }
        address := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
        bal, err := instance.BalanceOf(&bind.CallOpts{}, address) //&bind.CallOpts{} 只读调用，不需要打包交易上链，也不消耗 gas。
        if err != nil {
                log.Fatal(err)
        }
        name, err := instance.Name(&bind.CallOpts{})
        if err != nil {
                log.Fatal(err)
        }
        symbol, err := instance.Symbol(&bind.CallOpts{})
        if err != nil {
                log.Fatal(err)
        }
        decimals, err := instance.Decimals(&bind.CallOpts{})
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("name: %s\n", name)         // "name: Golem Network"
        fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
        fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
        fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"
        fbal := new(big.Float)
        fbal.SetString(bal.String())
        value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
        fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}