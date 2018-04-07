package storage

import (
	"encoding/json"
    "strconv"

	"github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
)

/// Struct for storing
type User struct {
	Username  string
	Password  string
	Balance   int
}

/// Struct for logging in
type UserQuery struct {
	Username string
	Password string
}

/// Store the user in the blockchain
func DeliverUser(app *StorageApplication, input []byte) {
	//input= { "MessageType":"User", "Username":"bijlar", "FirstName":"Adriaan", "LastName":"Bijl", "Password":"DnbBPCFu1ngp7DQFlNMrh3AKZE/VbdJ2J9TXTqPlraA=" }
	var data User
	json.Unmarshal(input, &data)

	app.Users[data.Username] = data
}

/// Check the user for validation errors for registration
func CheckUser(app *StorageApplication, input []byte) string {
	var data User
	err := json.Unmarshal(input, &data)

	// input= { "MessageType":"User", "Username":"bijlar", "Password":"DnbBPCFu1ngp7DQFlNMrh3AKZE/VbdJ2J9TXTqPlraA=" }
	if err != nil {
		return "Invalid input. JSON badly formatted."
	}

	if data.Username == "" {
		return "Invalid input. Username is empty."
	}

	if data.Password == "" {
		return "Invalid input. password is empty."
	}

	userAccount := app.Users[data.Username]
	if userAccount.Username != "" {
		return cmn.Fmt("Invalid input. User '%v' already exists.", data.Username)
	}

	return ""
}

func QueryUser(app *StorageApplication, identity []byte) types.ResponseQuery {
	// identity= { "Username":"ijpemar5", "Password": "1h93219hd9d1" }
	var inputdata UserQuery

	err := json.Unmarshal(identity, &inputdata)
	if err != nil {
		return types.ResponseQuery{Log: cmn.Fmt("Error occurred. Invalid input. JSON badly formatted in User Query. %v", err)}
	}

	userAccount := app.Users[inputdata.Username]

	 if userAccount.Username != "" && userAccount.Password == inputdata.Password {
		userInfo := "{ \"Username\":\"" + userAccount.Username + "\", \"Balance\":\"" + strconv.Itoa(userAccount.Balance) + "\" }"
		return types.ResponseQuery{Value: []byte(cmn.Fmt("%v", userInfo))}
	} else {
		return types.ResponseQuery{Log: "Error occurred. Invalid username/password combination."}
	}
}
