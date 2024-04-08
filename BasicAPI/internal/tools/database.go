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

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}
	// mockDB is a struct that implements the interface

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
