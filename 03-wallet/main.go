package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("Error to get a pvk", err)
	}

	pData := crypto.FromECDSA(pvk)

	fmt.Println(hexutil.Encode(pData))

	pubData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(pubData))
	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
