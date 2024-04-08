package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/hemanG-G/GO_MINI_PROJECTS/tree/main/BasicAPI/api"
	"github.com/hemanG-G/GO_MINI_PROJECTS/tree/main/BasicAPI/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {

	var params = api.CoinBalanceParam{} // get the Incoming request username
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query()) // store the incoming Username into the CoinBalanceParam struct

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase() //initialize new database
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username) // get token details from database using  username & middleware auth
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{ // Setting up Respose to be sent
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response) // encode Response into Json
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
