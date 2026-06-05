package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jetqin/goapi/api"
	"github.com/jetqin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params api.CoinBalanceParams
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error("Failed to decode query parameters: ", err)
		api.RequestErrorHander(w, err)
		return
	}

	database, err := tools.NewDatabaseInterface()
	if err != nil {
		log.Error("Failed to connect to database: ", err)
		api.InternalErrorHandler(w, err)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails, err = (*database).GetUserCoinDetails(params.Username)
	if err != nil {
		log.Error("Failed to fetch coin details: ", err)
		api.InternalErrorHandler(w, err)
		return
	}

	response := api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: (*tokenDetails).Coins,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Failed to encode response: ", err)
		api.InternalErrorHandler(w, err)
		return
	}
}
