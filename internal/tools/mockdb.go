package tools

import "time"

type mockdb struct{}

var mockUserDetails = map[string]LoginDetails{
	"karthik":{
		AuthToken: "ABC",
		name: "karthik",
	},
	"Dharanish":{
		AuthToken: "xyz",
		name: "Dharanish",
	},
}

var mockCoinDetails = map[string] CoinDetails{
	"karthik":{
		Coins: 23,
		username: "karthik",
	},
	"Dharanish":{
		Coins: 34,
		username: "Dharanish",
	},
}

func (d *mockdb) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientDetails LoginDetails
	clientDetails, ok := mockUserDetails[username]

	if !ok{
		return  nil
	}

	return &clientDetails
}

func (d *mockdb) GetCoinDetails(username string) *CoinDetails{
	time.Sleep(time.Second * 1)

	var coinDetails CoinDetails

	coinDetails, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &coinDetails
}

func (d *mockdb) SetupDatabase() error{
	return nil
}