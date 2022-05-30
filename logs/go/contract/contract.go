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
type Log struct {
	Log string `json:"log"`
}

// the business object
type User struct {
	ID string `json:"id"`
	User string `json:"user"`
	Logs []Log `json:"logs"`
}

// create a new voting record
func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, id string, user string) error {
	user := User{
		ID: id,
		User: user,
		Logs: []Log{},
	}

	//create readable object for the database
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, userJSON)
} 

// update a voting record, For
func (s *SmartContract) Log(ctx contractapi.TransactionContextInterface, id string, log string) error {
	iuserJSON, err := ctx.GetStub().GetState(id)
	var user User
	json.Unmarshal(iuserJSON, &user)

	user.Logs = append(user.Logs, Log{Log: log})

	//create readable object for the database
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, userJSON)
} 

// get one vote record
func (s *SmartContract) GetUser(ctx contractapi.TransactionContextInterface, id string) (*User, error) {
	userJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}

	var user User
	json.Unmarshal(userJSON, &user)

	return &user, nil
} 

// get all the vote records
func (s *SmartContract) GetAllUsers(ctx contractapi.TransactionContextInterface) ([]*User, error) {
	result, err := ctx.GetStub().GetStateByRange("","")
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var users []*User
	for result.HasNext()  {
		usersJSON, err := result.Next()
		if err != nil {
			return nil, err
		}

		var user User
		json.Unmarshal(userJSON.Value, &user)
		users = append(users, &user)
	}

	return users, nil
} 