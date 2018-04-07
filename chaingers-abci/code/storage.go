package storage

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
)

/// Struct for storing
type StorageApplication struct {
	types.BaseApplication

	Users             map[string]User
	Transactions      map[string][]Transaction
}

// Initialisation
func NewStorageApplication() *StorageApplication {
	return &StorageApplication{Users: make(map[string]User), Transactions: make(map[string][]Transaction)}
}

// Storing data in the blockchain
func (app *StorageApplication) DeliverTx(tx []byte) types.Result {
	fmt.Println("Entering DeliverTx...")

	input := string(tx[:])

	// Map json to string array
	var data map[string]interface{}
	json.Unmarshal(tx[:], &data)

	inputBytes := bytes.Trim([]byte(input), string([]byte{0}))
	if (data["MessageType"] == "User") {
		// Store User
		DeliverUser(app, inputBytes)
	} else if (data["MessageType"] == "Transaction") {
		// Store transaction
		DeliverTransaction(app, inputBytes)
	}

	return types.OK
}

// Validation of data before storing in blockchain
func (app *StorageApplication) CheckTx(tx []byte) types.Result {
	fmt.Println("Entering CheckTx...")

	input := string(tx[:])

	// Map json naar string array
	var data map[string]interface{}
	err := json.Unmarshal(tx[:], &data)

	if err != nil {
		return types.ErrBaseInvalidInput.SetLog(cmn.Fmt("Error occurred. JSON badly formatted. %v", err))
	}

	inputBytes := bytes.Trim([]byte(input), string([]byte{0}))
	if (data["MessageType"] == "User") {
		//Check user for validation errors
		error := CheckUser(app, inputBytes)
		if error != "" {
			return types.ErrBaseInvalidInput.SetLog(error)
		} else {
			//No errors found
			return types.OK
		}
	} else if (data["MessageType"] == "Transaction") {
		error := CheckTransaction(app, inputBytes)
		if error != "" {
			return types.ErrBaseInvalidInput.SetLog(error)
		} else {
			//No errors found
			return types.OK
		}
	} else {
		return types.ErrBaseInvalidInput.SetLog("Unknown MessageType")
	}
}

// What to do when the data is about to be committed in the blockchain
func (app *StorageApplication) Commit() types.Result {
	fmt.Println("Entering Commit...")
	return types.OK
}

// Retrieving general data
func (app *StorageApplication) Info(req types.RequestInfo) types.ResponseInfo {
	fmt.Println("Entering Info...")
	return types.ResponseInfo{Data: cmn.Fmt("{ users: %v }", len(app.Users))}
}

// Query the blockchain for data
func (app *StorageApplication) Query(reqQuery types.RequestQuery) types.ResponseQuery {
	fmt.Println("Entering Query...")

	// No input values found, then return error
	if reqQuery.Data == nil || len(reqQuery.Data) == 0 {
		return types.ResponseQuery{Log: "Error occurred. 'Data' does not exist or is empty."}
	}

	input := string(reqQuery.Data)

	// Map json naar string array
	var data map[string]interface{}
	err := json.Unmarshal(reqQuery.Data, &data)
	if err != nil {
		return types.ResponseQuery{Log: cmn.Fmt("Error occurred. JSON badly formatted in Query. %v", err)}
	}

	inputBytes := bytes.Trim([]byte(input), string([]byte{0}))
	if reqQuery.Path == "User" {
		// Find user
		return QueryUser(app, inputBytes)
	} else if reqQuery.Path == "Transaction" {
		// Find transaction(s)
		return QueryTransaction(app, inputBytes);
	}

	// No type of data defined in path, so error is given
	return types.ResponseQuery{Log: cmn.Fmt("Error occurred. Path '%v' is not valid.", reqQuery.Path)}
}
