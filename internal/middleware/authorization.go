package middleware

import (
	"errors"
	"net/http"

	"github.com/jetqin/goapi/api"
	"github.com/jetqin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		var err error
		if username == "" || token == "" {
			log.Warnf("Missing username or token. Username: %s, Token: %s", username, token)
			log.Error(UnAuthorizedError)
			api.RequestErrorHander(w, UnAuthorizedError)
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabaseInterface()
		if err != nil {
			log.Error("Failed to connect to database: ", err)
			api.InternalErrorHandler(w, err)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails, err = (*database).GetUserLoginDetails(username)
		if (loginDetails == nil) || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHander(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
