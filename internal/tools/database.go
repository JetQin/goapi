package tools

import (
	"sync"

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

var (
	globalDB *sqliteDB
	initOnce sync.Once
)

func NewDatabaseInterface() (*DatabaseInterface, error) {
	log.Info("Connecting to database...")

	var initErr error
	initOnce.Do(func() {
		db := &sqliteDB{}
		initErr = db.SetupDatabase()
		if initErr != nil {
			log.Error("Failed to set up sqlite database: ", initErr)
			return
		}
		globalDB = db
	})

	if initErr != nil {
		return nil, initErr
	}

	var database DatabaseInterface = globalDB
	return &database, nil
}
