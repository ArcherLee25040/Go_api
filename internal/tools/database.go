package tools

import "github.com/opentracing/opentracing-go/log"

//Database collections
type LoginDetail struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDatails(username string) *LoginDetail
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()

	if err != nil {
		log.Error(err)
		return nil, err

	}

	return &database, nil
}
