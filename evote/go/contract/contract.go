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
type Vote struct {
	Node string `json:"node"`
	Vote string `json:"vote"`
}

// the business object
type Poll struct {
	ID string `json:"id"`
	Creator string `json:"creator"`
	Status bool `json:"status"`
	Votes []Vote `json:"votes"`
}

// create a new voting record
func (s *SmartContract) CreatePoll(ctx contractapi.TransactionContextInterface, id string, creator string) error {
	poll := Poll{
		ID: id,
		Creator: creator,
		Status: true,
		Votes: []Vote{},
	}

	//create readable object for the database
	pollJSON, err := json.Marshal(poll)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, pollJSON)
} 

// update a voting record, For
func (s *SmartContract) Approve(ctx contractapi.TransactionContextInterface, id string, node string) error {
	ipollJSON, err := ctx.GetStub().GetState(id)
	var poll Poll
	json.Unmarshal(ipollJSON, &poll)

	for i, vote := range poll.Votes {
		if (vote.Node == "Node") {
			poll.Votes[i].Vote = "1"
			//create readable object for the database
			pollJSON, err := json.Marshal(poll)
			if err != nil {
				return err
			}
			return ctx.GetStub().PutState(id, pollJSON)
		}
    }

	poll.Votes = append(poll.Votes, Vote{Node: node, Vote: "1"})

	//create readable object for the database
	pollJSON, err := json.Marshal(poll)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, pollJSON)
} 

// update a voting record, Against
func (s *SmartContract) Decline(ctx contractapi.TransactionContextInterface, id string, node string) error {
	ipollJSON, err := ctx.GetStub().GetState(id)
	var poll Poll
	json.Unmarshal(ipollJSON, &poll)

    for i, vote := range poll.Votes {
		if (vote.Node == "Node") {
			poll.Votes[i].Vote = "0"
			//create readable object for the database
			pollJSON, err := json.Marshal(poll)
			if err != nil {
				return err
			}
			return ctx.GetStub().PutState(id, pollJSON)
		}
    }

	poll.Votes = append(poll.Votes, Vote{Node: node, Vote: "1"})

	//create readable object for the database
	pollJSON, err := json.Marshal(poll)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, pollJSON)
} 

// close voting session
func (s *SmartContract) ClosePoll(ctx contractapi.TransactionContextInterface, id string) error {
	poll, err := s.GetPoll(ctx, id)
	if err != nil {
		return err
	}

	poll.Status = false
	
	pollJSON, err := json.Marshal(poll)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, pollJSON)
} 

// get one vote record
func (s *SmartContract) GetPoll(ctx contractapi.TransactionContextInterface, id string) (*Poll, error) {
	pollJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}

	var poll Poll
	json.Unmarshal(pollJSON, &poll)

	return &poll, nil
} 

// get all the vote records
func (s *SmartContract) GetAllPolls(ctx contractapi.TransactionContextInterface) ([]*Poll, error) {
	result, err := ctx.GetStub().GetStateByRange("","")
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var polls []*Poll
	for result.HasNext()  {
		pollJSON, err := result.Next()
		if err != nil {
			return nil, err
		}

		var poll Poll
		json.Unmarshal(pollJSON.Value, &poll)
		polls = append(polls, &poll)
	}

	return polls, nil
} 