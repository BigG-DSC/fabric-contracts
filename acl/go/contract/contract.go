package contract

import (
	"encoding/json"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// the business object
type AccessControlElem struct {
	ID string `json:"id"`
	Creator string `json:"creator"`
	CID string `json:"cid"`
}

// creat a new record in the Access Control List
func (s *SmartContract) Store(ctx contractapi.TransactionContextInterface, id string, creator string, cid string) error {
	aclelem := AccessControlElem{
		ID: id,
		Creator: creator,
		CID: cid,
	}

	//create readable object for the database
	aclelemJSON, err := json.Marshal(aclelem)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, aclelemJSON)
} 

// check if you are part of the Access Control List
func (s *SmartContract) Check(ctx contractapi.TransactionContextInterface, id string) (*AccessControlElem, error) {
	aclelemJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}

	var aclelem AccessControlElem
	json.Unmarshal(aclelemJSON, &aclelem)

	return &aclelem, nil
} 