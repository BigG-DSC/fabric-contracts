package main

import (
	"log"
	"acl/go/contract"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	aclChaincode, err := contractapi.NewChaincode(&contract.SmartContract{})
	if err != nil {
		log.Panicf("Error: %v", err)
	}
	if err := aclChaincode.Start(); err != nil {
		log.Panicf("Error: %v", err)
	}
}
