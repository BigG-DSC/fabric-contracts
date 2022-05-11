package main

import (
	"log"
	"evote/go/contract"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	evoteChaincode, err := contractapi.NewChaincode(&contract.SmartContract{})
	if err != nil {
		log.Panicf("Error: %v", err)
	}
	if err := evoteChaincode.Start(); err != nil {
		log.Panicf("Error: %v", err)
	}
}
