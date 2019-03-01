package main

import (
	"crypto/ecdsa"
	"fmt"
	//"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//client, err := ethclient.Dial("https://mainnet.infura.io") //non-production only
	//client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc") //for when running a local geth client
	client, err := ethclient.Dial("http://localhost:8545") //local testing only

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection")
	_ = client // we'll use this in the upcoming sections

	//address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	//fmt.Println(address.Hex())

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privateKey)
	privateKeyBytes := crypto.FromECDSA(privateKey) //convert the key to bytes
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	fmt.Println(publicKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
}
