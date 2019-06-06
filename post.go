package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/Task/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

const key  = `{"address":"a8112ac2f02fa71f737929d18671b72e8609b78d","crypto":{"cipher":"aes-128-ctr","ciphertext":"903f9828593c5b734557b8a9d2d14b19a43da10f12cbd532fa3d231bf13eec0f","cipherparams":{"iv":"25160f43bc0c48b2ba29e51ae304680f"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d707e45ca598f3bd5e71e3d3acd3ec656d5adbf35fa308ab6e4e88a1e643a7c9"},"mac":"113f69f6cd267e242ddb0e640bd195a12637c128a5dc83d77a266fe22176bf8b"},"id":"85285b23-2ea6-4a4d-958b-84d6a06d1e23","version":3}`

func main(){
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/0e659bef259e47fc95c902d9fdf30767")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "agtl1234$")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	
  contract, err :=contracts.NewSampel(common.HexToAddress("0xab33B1F430E4F6c9872780CE3769830390C6C1C5"), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	var p,_ = contract.SetUser(&bind.TransactOpts{
		From:auth.From,
		Signer:auth.Signer,
		Value: nil,
	}, "qw")


	if  p != nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}