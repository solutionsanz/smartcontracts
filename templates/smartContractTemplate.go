/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the @CONTRACT_NAME_LC@ structure, with 4 properties.  Structure tags are used by encoding/json library
type @CONTRACT_NAME@ struct {	
	@NEW_CONTRACT_STRUCTURE@
}

/*
 * The Init method is called when the Smart Contract "fab@CONTRACT_NAME_LC@" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fab@CONTRACT_NAME_LC@"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "query@CONTRACT_NAME@" {
		return s.query@CONTRACT_NAME@(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "create@CONTRACT_NAME@" {
		return s.create@CONTRACT_NAME@(APIstub, args)
	} else if function == "queryAll@CONTRACT_NAME@s" {
		return s.queryAll@CONTRACT_NAME@s(APIstub)
	} else if function == "record@CONTRACT_NAME@Action" {
		return s.record@CONTRACT_NAME@Action(APIstub, args)
	} else if function == "getHistoryForRecord" {
        return s.getHistoryForRecord(APIstub, args)
    }


	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) query@CONTRACT_NAME@(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	@CONTRACT_NAME_LC@AsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(@CONTRACT_NAME_LC@AsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	@CONTRACT_NAME_LC@s := []@CONTRACT_NAME@{
		@NEW_INITLEDGER_CONTRACT_ITEM@
	}

	i := 0
	for i < len(@CONTRACT_NAME_LC@s) {
		fmt.Println("i is ", i)
		@CONTRACT_NAME_LC@AsBytes, _ := json.Marshal(@CONTRACT_NAME_LC@s[i])
		APIstub.PutState("@CONTRACT_NAME_UC@"+strconv.Itoa(i), @CONTRACT_NAME_LC@AsBytes)
		fmt.Println("Added", @CONTRACT_NAME_LC@s[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) create@CONTRACT_NAME@(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != @NUMBER_OF_CONTRACT_ITEMS@ {
		return shim.Error("Incorrect number of arguments. Expecting @NUMBER_OF_CONTRACT_ITEMS@")
	}

	var @CONTRACT_NAME_LC@ = @CONTRACT_NAME@@CONST_CREATE_NEW_RECORD_LINE@

	@CONTRACT_NAME_LC@AsBytes, _ := json.Marshal(@CONTRACT_NAME_LC@)
	APIstub.PutState(args[0], @CONTRACT_NAME_LC@AsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAll@CONTRACT_NAME@s(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "@CONTRACT_NAME_UC@0"
	endKey := "@CONTRACT_NAME_UC@999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAll@CONTRACT_NAME@s:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) record@CONTRACT_NAME@Action(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != @NUMBER_OF_CONTRACT_ITEMS@ {
		return shim.Error("Incorrect number of arguments. Expecting @NUMBER_OF_CONTRACT_ITEMS@")
	}

	@CONTRACT_NAME_LC@AsBytes, _ := APIstub.GetState(args[0])
	@CONTRACT_NAME_LC@ := @CONTRACT_NAME@{}

	json.Unmarshal(@CONTRACT_NAME_LC@AsBytes, &@CONTRACT_NAME_LC@)
	@CONTRACT_ITEM_RECORD@

	@CONTRACT_NAME_LC@AsBytes, _ = json.Marshal(@CONTRACT_NAME_LC@)
	APIstub.PutState(args[0], @CONTRACT_NAME_LC@AsBytes)

	return shim.Success(nil)
}


// ===========================================================================================
// getHistoryForRecord returns the histotical state transitions for a given key of a record
// ===========================================================================================
func  (s *SmartContract) getHistoryForRecord(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    
    if len(args) < 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    recordKey := args[0]

    fmt.Printf("- start getHistoryForRecord: %s\n", recordKey)
    
    resultsIterator, err := APIstub.GetHistoryForKey(recordKey)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    // buffer is a JSON array containing historic values for the key/value pair
    var buffer bytes.Buffer
    buffer.WriteString("[")

    bArrayMemberAlreadyWritten := false
    for resultsIterator.HasNext() {
        response, err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        // Add a comma before array members, suppress it for the first array member
        if bArrayMemberAlreadyWritten == true {
            buffer.WriteString(",")
        }
        buffer.WriteString("{\"TxId\":")
        buffer.WriteString("\"")
        buffer.WriteString(response.TxId)
        buffer.WriteString("\"")

        buffer.WriteString(", \"Value\":")
        // if it was a delete operation on given key, then we need to set the
        //corresponding value null. Else, we will write the response.Value
        //as-is (as the Value itself a JSON vehiclePart)
        if response.IsDelete {
            buffer.WriteString("null")
        } else {
            buffer.WriteString(string(response.Value))
        }

        buffer.WriteString(", \"Timestamp\":")
        buffer.WriteString("\"")
        buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
        buffer.WriteString("\"")

        buffer.WriteString(", \"IsDelete\":")
        buffer.WriteString("\"")
        buffer.WriteString(strconv.FormatBool(response.IsDelete))
        buffer.WriteString("\"")

        buffer.WriteString("}")
        bArrayMemberAlreadyWritten = true
    }
    buffer.WriteString("]")

    fmt.Printf("- getHistoryForRecord returning:\n%s\n", buffer.String())

    return shim.Success(buffer.Bytes())
}


// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
