package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username  string
	AuthToken string
}

type CoinDetails struct {
	Username string
	Coins    int64
}

type DatabaseInterface interface {
	// In a real implementation, this would likely be a connection pool or client for a database.
	GetUserLoginDetails(username string) (*LoginDetails, error)
	GetUserCoinDetails(username string) (*CoinDetails, error)
	SetupDatabase() error
}

func NewDatabaseInterface() (*DatabaseInterface, error) {
	// In a real implementation, this would establish a connection to the database and return an interface to interact with it.
	log.Info("Connecting to database...")

	var database DatabaseInterface = &mockDB{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error("Failed to set up database: ", err)
		return nil, err
	}

	return &database, nil
}
