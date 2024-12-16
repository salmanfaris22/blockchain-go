package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Url = "https://mainnet.infura.io/v3/8d8280da4b8f410c9964968301452e53"
var ganacheUrl = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), Url)
	if err != nil {
		log.Fatal("Error to get a client", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("Error to get a block", err)
	}
	fmt.Println(block.Number())

	addr := "0x577597fc5c94F5bc5FF203F3414A50ED403f0782"
	address := common.HexToAddress(addr)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("Error to get a balance", err)
	}

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	balanceEther := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("balanceEther is", balanceEther)
}
