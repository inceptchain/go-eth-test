package main

import (
	"crypto/ecdsa"
	"fmt"
	"context"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	//"math/big"

	token "./token"
)

func main() {
	//client, err := ethclient.Dial("https://ropsten.infura.io/NAZ50S21bUu0MG5BSxMx") //non-production only
	//client, err := ethclient.Dial("wss://ropsten.infura.io/v3/NAZ50S21bUu0MG5BSxMx") //non-production only 
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/db218500773c4d6cbe2844e0d40673b5") //non-production only
	//client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc") //for when running a local geth client
	//client, err := ethclient.Dial("http://localhost:8545") //local testing only

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection")

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String())

	tokenAddress := common.HexToAddress("0xe8dd269b096382cdcab10ab5b6c188a2d05a9300")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)

	privateKey, err := crypto.HexToECDSA("B8FDDA31565B2A620BF391EDD3F844A9F4FDB5A3411DE80BD223EE0EA5DC7C47")
	//privateKey, err := crypto.HexToECDSA("B71F491BF9DAC623AF2B87B66B8325B5058AAB27AB3AEA133DF240F74E671EE1")//vlad pkey
	if err != nil {
  		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
  		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
  		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
  		log.Fatal(err)
	}
	/*
		struct Structure {
			uint number
			string phrase
		}
	*/
	//hash something into bytes32 and put it into mapping(bytes32 => bool)
	//put something into mapping(address => Structure)

	fmt.Println(nonce)
	fmt.Println(fromAddress)//how do I make this into a normal address?
	fmt.Println(address)//how do I make this into a normal address?
	fmt.Println(gasPrice)

	// value := big.NewInt(1000000000000000000)
	// gasLimit := uint64(21000)
	// //toAddress := common.HexToAddress("0x8608e7313a107386D434d2Dbf68575315EC40ef8")
	// toAddress := common.HexToAddress("0x7A150f31819a002e1322721F9B0bEbEe713edaA9")
	// var data []byte
	// tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// chainId, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
