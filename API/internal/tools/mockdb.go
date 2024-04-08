package tools

import (
	"time"
)

// mock implementation of a database
type mockDB struct{}

// we define methods outside structs -> called reciever functions
// func (s *AssociatedStruct) MethodName(params.) ...

var mockLoginDetails = map[string]LoginDetails{ // mock implementation of a database

	"alex": {
		AuthToken: "password",
		Username:  "hemang",
	},
	"jason": {
		AuthToken: "password",
		Username:  "aman",
	},
	"marie": {
		AuthToken: "password",
		Username:  "karthik",
	},
}

var mockCoinDetails = map[string]CoinDetails{ // mock implementation of a database

	"alex": {
		Coins:    100,
		Username: "hemang",
	},
	"jason": {
		Coins:    200,
		Username: "aman",
	},
	"marie": {
		Coins:    300,
		Username: "karthik",
	},
}

// func (s *AssociatedStruct) MethodName(params.) ...
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails { //mockDB represents the instance of the DATABASE we want to access
	// Simulate DB call by adding delay
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok { // checking clientData filled or not
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails { //mockDB represents the instance of the DATABASE we want to access
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
