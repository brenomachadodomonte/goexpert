package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
}

func main() {
	account := Account{Number: 1, Balance: 200}
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(account)
	if err != nil {
		println(err)
	}

	plainJson := []byte(`{"number":2,"balance":300}`)
	var myAccount Account
	err = json.Unmarshal(plainJson, &myAccount)
	if err != nil {
		println(err)
	}
	println(myAccount.Balance)

}
