package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct{
	AuthToken string
	name string
}

type CoinDetails struct{
	Coins int
	username string
}

type DatabaseInterface interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetCoinDetails(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface = &mockdb{}

	var error error = database.SetupDatabase()

	if error != nil{
		log.Error(error)
		return nil, error
	}

	return &database, nil

}