/********************************************************************************
*********************************************************************************
*********************************************************************************
MIT License

Copyright (c) 2022 Gioele Bigini

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
********************************************************************************
********************************************************************************
********************************************************************************/

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