package middleware

import (
	"fmt"
	"net/http"

	"github.com/kaizakin/goapi/api"
	"github.com/kaizakin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var unAuthorizedError = fmt.Errorf("the user is unauthorized")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		
		if username == ""{
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}
		
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil{
			api.RequestInternalHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if(loginDetails == nil ||( (*loginDetails).AuthToken != token)){
			log.Error(unAuthorizedError)
			api.RequestInternalHandler(w)
			return
		}

		next.ServeHTTP(w,r)
	})
}