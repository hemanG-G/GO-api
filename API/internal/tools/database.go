package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collections
// db return rypes
type LoginDetails struct {
	AuthToken string
	Username  string
}
type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface { // interacting with the DB
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// initialize new database
func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{} // mockDB is an instance/implementation of a database created
	// mockDB is a struct that implements the interface
	//every time new mockDB setup

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
