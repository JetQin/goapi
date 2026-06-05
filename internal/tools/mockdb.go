package tools

import (
	"time"
)

type mockDB struct {}

type mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"bob": {
		AuthToken	: "456DEF",
		Username:  "bob",
	}
}

type mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"bob": {
		Coins: 50,
		Username: "bob",
	}
}

func (db *mockDB) GetUserLoginDetails(username string) (*LoginDetails, error) {
	time.Sleep(100 * time.Millisecond) // Simulate database latency
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil, fmt.Errorf("user not found: %s", username)
	}
	return &clientData, nil
}

func (db *mockDB) GetUserCoinDetails(username string) (*CoinDetails, error) {
	time.Sleep(100 * time.Millisecond) // Simulate database latency
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil, fmt.Errorf("user not found: %s", username)
	}
	return &clientData, nil
}

func (db *mockDB) SetupDatabase() error {
	// In a real implementation, this would set up the database schema and any necessary tables.
	// For the mock, we can just return nil to indicate success.
	return nil
}
	