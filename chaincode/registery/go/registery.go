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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the hash structure, with 2 properties.  Structure tags are used by encoding/json library
type R_Hash struct {
	Hash  string `json:"hash"`
	District string `json:"district"`
}

/*
 * The Init method is called when the Smart Contract "registery" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "registery"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryHash" {
		return s.queryHash(APIstub, args)
	 } else if function == "initLedger" {
	 	return s.initLedger(APIstub)
	 } else if function == "addHash" {
		return s.addHash(APIstub, args)
	} else if function == "queryAllHashes" {
		return s.queryAllHashes(APIstub)
	} else if function == "changeHash" {
		return s.changeHash(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryHash(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	hashAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(hashAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	Hashes := []R_Hash{
		R_Hash{District: "01",	Hash: "7cb6fa91c124913f7a75e3153339234f"},
		R_Hash{District: "02",	Hash: "a2a551a6458a8de22446cc76d639a9e9"},
		R_Hash{District: "03",	Hash: "0cc175b9c0f1b6a831c399e269772661"},
		R_Hash{District: "04",	Hash: "f016441d00c16c9b912d05e9d81d894d"},
		R_Hash{District: "05",	Hash: "755f85c2723bb39381c7379a604160d8"},
		R_Hash{District: "06",	Hash: "1a699ad5e06aa8a6db3bcf9cfb2f00f2"},
	}

	i := 0
	for i < len(Hashes) {
		fmt.Println("i is ", i)
		hashAsBytes, _ := json.Marshal(Hashes[i])
		APIstub.PutState("R_Hash"+strconv.Itoa(i), hashAsBytes)
		fmt.Println("Added", Hashes[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) addHash(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	var hash = R_Hash{District: args[1], Hash: args[2]}

	hashAsBytes, _ := json.Marshal(hash)
	APIstub.PutState(args[0], hashAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllHashes(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "R_Hash0"
	endKey := "R_Hash999"

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

	fmt.Printf("- queryAllHashes:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeHash(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	hashAsBytes, _ := APIstub.GetState(args[0])
	hash := R_Hash{}

	json.Unmarshal(hashAsBytes, &hash)
	hash.District = args[1]

	hashAsBytes, _ = json.Marshal(hash)
	APIstub.PutState(args[0], hashAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
