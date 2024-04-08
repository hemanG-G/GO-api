package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hemanG-G/GO_MINI_PROJECTS/tree/main/BasicAPI/api"
	"github.com/hemanG-G/GO_MINI_PROJECTS/tree/main/BasicAPI/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New(fmt.Sprintf("Invalid username or token."))

// authorization middleware called before /account/coins route

func Authorization(next http.Handler) http.Handler { // 'next' is the next HTTP handler after this middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// The purpose of returning a handler function is to wrap the logic of the middleware
		// around the execution of the next handler.

		//middleware logic
		var username string = r.URL.Query().Get("username") // get username from URL
		var token = r.Header.Get("Authorization")           // get auth token from header
		var err error

		if username == "" {
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase() // initialize database
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username) // get login details from DB , to match with incoming details

		// auth incoming details
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// next handler
		next.ServeHTTP(w, r) //continue onto from this middleware to next handler

	})
}
