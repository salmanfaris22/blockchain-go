package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "1234"
	// a, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal("Error to get a NewAccount", err)
	// }
	// fmt.Println(a.Address)
	b, err := ioutil.ReadFile("./wallet/UTC--2024-12-16T05-47-03.028541000Z--b6227aaa8e0354daf380acfd97697e9c71697a90")
	if err != nil {
		log.Fatal("Error to get a ioutil", err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal("Error to get a keystore", err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println(hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println(hexutil.Encode(pData))
	fmt.Println(crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}
