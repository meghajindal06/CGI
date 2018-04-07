package storage

import (
	"encoding/json"
    "strconv"

	"github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
)

/// Struct for storing
type Transaction struct {
	Capacity	float64
	Fee			float64
}

/// Struct for input
type TransactionInput struct {
	Sender		string
	Receiver	string
	Capacity	string
	Fee			string
}

/// Struct for searching
type TransactionQuery struct {
	Username	string
}

/// Store the transaction in the blockchain
func DeliverTransaction(app *StorageApplication, input []byte) {
	//input= { "MessageType":"Transaction", "Sender":"bijlar", "Receiver": "ijpemar", "Capacity": "120.0", "Fee": "5.52" }
	var data TransactionInput
	json.Unmarshal(input, &data)

	var newData Transaction
	newData.Capacity, _ = strconv.ParseFloat(data.Capacity, 64);
	newData.Fee, _ = strconv.ParseFloat(data.Fee, 64);

	app.Transactions[data.Sender] = append(app.Transactions[data.Sender], newData)
	app.Transactions[data.Receiver] = append(app.Transactions[data.Receiver], newData)
}

/// Check the user for validation errors for registration
func CheckTransaction(app *StorageApplication, input []byte) string {
	//input= { "MessageType":"Transaction", "Sender":"bijlar", "Receiver": "ijpemar", "Capacity": "120.0", "Fee": "5.52" }v
	var data TransactionInput
	err := json.Unmarshal(input, &data)

	if err != nil {
		return "Invalid input. JSON badly formatted."
	}

	if data.Sender == "" {
		return "Invalid input. Sender is empty."
	}

	if data.Receiver == "" {
		return "Invalid input. Receiver is empty."
	}

	if data.Capacity == "" {
		return "Invalid input. Capacity is empty."
	}

	if data.Fee == "" {
		return "Invalid input. Fee is empty."
	}

	return ""
}

func QueryTransaction(app *StorageApplication, identity []byte) types.ResponseQuery {
	// identity= { "Username":"ijpemar" }
	var inputdata TransactionQuery
	err := json.Unmarshal(identity, &inputdata)
	if err != nil {
		return types.ResponseQuery{Log: cmn.Fmt("Error occurred. Invalid input. JSON badly formatted in TransactionQuery. %v", err)}
	}

	transactions := app.Transactions[inputdata.Username]

	if transactions != nil {
		jsonResult, err := json.Marshal(transactions);

		if (err == nil) {
			return types.ResponseQuery{Value: []byte(cmn.Fmt("%v", string(jsonResult)))};
		} else {
			return types.ResponseQuery{Log: cmn.Fmt("Error occurred. Could not Marshall JSON source. %v", err)};
		}
	} else {
		return types.ResponseQuery{Log: "Error occurred. Invalid username."}
	}
}
