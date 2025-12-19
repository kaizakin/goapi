package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/kaizakin/goapi/api"
	"github.com/kaizakin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request){
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil{
		log.Error(err)
		api.RequestInternalHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if database == nil{
		log.Error(err)
		api.RequestInternalHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetCoinDetails(params.Username)

	if tokenDetails == nil{
		log.Error(err)
		api.RequestInternalHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: int64((*tokenDetails).Coins),
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil{
		log.Error(err)
		api.RequestInternalHandler(w)
		return 
	}
}