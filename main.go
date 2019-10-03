package main

import (
	"crypto/ecdsa"
	"fmt"
	//"context"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	//"math/big"

	//token "./token"
	mapping "./mapping" 
)

func main() {
	//client, err := ethclient.Dial("https://ropsten.infura.io/NAZ50S21bUu0MG5BSxMx") //non-production only
	//client, err := ethclient.Dial("wss://ropsten.infura.io/v3/NAZ50S21bUu0MG5BSxMx") //non-production only 
	//client, err := ethclient.Dial("https://ropsten.infura.io/v3/db218500773c4d6cbe2844e0d40673b5") //non-production only
	//client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc") //for when running a local geth client
	client, err := ethclient.Dial("http://localhost:8545") //local testing only

	if err != nil {
		log.Fatalf("client err: %v", err)
	}

	fmt.Println("We have a connection")
	
	//exit infura or local connection
	defer client.Close() //necessary or not?

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("header error: ",err)
	// }

	// fmt.Println(header.Number.String())

	mappingAddress := common.HexToAddress("0x4d86803ad13dfd01da8e745652775eaac7dcdfdf")//
	//tokenAddress := common.HexToAddress("0xe8dd269b096382cdcab10ab5b6c188a2d05a9300")//ganache-cli changes every restart of ganache
	//tokenAddress := common.HexToAddress("0x22E6a5341d8E92c08d2A565F636d95C3D7449238")//testnet
	//instance, err := token.NewToken(tokenAddress, client)
	instance, err := mapping.NewMapping(mappingAddress, client)
	if err != nil {
		log.Fatalf("instance err: %v", err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("name err: %v", err)
	}

	fmt.Printf("name: %s\n", name)

	privateKey, err := crypto.HexToECDSA("c1b3b8dac7cb70e62d2dcf313c58ddb27d175d83a76f28c7fa7fd86c9db4e3a1")//ganache-cli pkey
	//privateKey, err := crypto.HexToECDSA("B8FDDA31565B2A620BF391EDD3F844A9F4FDB5A3411DE80BD223EE0EA5DC7C47")//0x7A150f31819a002e1322721F9B0bEbEe713edaA9
	//privateKey, err := crypto.HexToECDSA("B71F491BF9DAC623AF2B87B66B8325B5058AAB27AB3AEA133DF240F74E671EE1")//vlad pkey
	if err != nil {
  		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
  		log.Fatal("error casting public key to ECDSA")
	}

	addressCheck := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
 // 			log.Fatalf("nonce error: %v",err)
	// }

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
 //  		log.Fatal(err)
	// }

	// auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(2))
	// auth.Value = big.NewInt(0)
	// auth.GasLimit = uint64(3000000)
	// auth.GasPrice = gasPrice
	//auth.Context = nil 

	addrs := common.HexToAddress("0x427a7eacc6af96571c458a1da7093c127887b857") // account(1) ganache-cli
	

	/*
		Transaction setTestData
	*/
	
	// num := big.NewInt(11)
	// phrs := "Cool Beans"

	// tx, err := instance.SetTestData(auth, addrs, num, phrs)
	// if err != nil {
	// 	log.Fatalf("tx error: %v",err)
	// }
	// fmt.Printf("test data tx sent: %s\n", tx.Hash().Hex())


	/*
		Transaction setBytesBool
	*/
	
	 value := [32]byte{}
	 copy(value[:], []byte("whattt"))
	
	// tx, err := instance.SetBytesBool(auth, value, true)
	// if err != nil {
	// 	log.Fatalf("bytes tx error: %v",err)
	// }
	// fmt.Printf("tx sent: %s\n", tx.Hash().Hex())


	result1, err := instance.AddressTest(nil, addrs)
	if err != nil {
		log.Fatalf("Number error: %d", err)
	}
	fmt.Printf("Should be number: %v\n", result1.Number)


	result2, err := instance.AddressTest(nil, addrs)
	if err != nil {
		log.Fatalf("Phrase error: %v",err)
	}
	fmt.Printf("Should be phrase: %v\n", result2.Phrase)


	result3, err := instance.GetOwner(nil)
	if err != nil {
		log.Fatalf("Get owner add: %v", err)
	}
	fmt.Printf("Owner address: %v\n", result3.Hex())

	result4, err := instance.BytesTest(nil, value)
	if err != nil {
		log.Fatalf("Bytes test error: %v", err)
	}
	fmt.Printf("True or False: %v\n", result4)

	/*
		Deploy Contract
	*/
	// name := "Hello World"
	// address, tx, instance, err := mapping.DeployMapping(auth, client, name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(address.Hex())
	// fmt.Println(tx.Hash().Hex())
	// _ = instance

	//fmt.Println(nonce)
	// fmt.Println(fromAddress) 
	fmt.Println(addressCheck)
	// fmt.Println(gasPrice)
	
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
