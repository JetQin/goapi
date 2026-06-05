package tools

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type sqliteDB struct {
	db *sql.DB
}

func (s *sqliteDB) dbPath() string {
	// place DB in project root next to go.mod
	wd, err := os.Getwd()
	if err != nil {
		return "goapi.db"
	}
	return filepath.Join(wd, "goapi.db")
}

func (s *sqliteDB) SetupDatabase() error {
	path := s.dbPath()
	// ensure directory exists
	_ = os.MkdirAll(filepath.Dir(path), 0755)

	// open sqlite file
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return fmt.Errorf("open sqlite: %w", err)
	}
	// set sensible timeouts
	db.SetConnMaxLifetime(time.Minute * 5)
	s.db = db

	// create table
	create := `CREATE TABLE IF NOT EXISTS users (
	username TEXT PRIMARY KEY,
	authtoken TEXT,
	coins INTEGER
);`
	_, err = s.db.Exec(create)
	if err != nil {
		return fmt.Errorf("create table: %w", err)
	}

	// seed data if not exists
	seed := `INSERT OR IGNORE INTO users(username, authtoken, coins) VALUES
	('alex', '123ABC', 100),
	('bob', '456DEF', 50);`
	_, err = s.db.Exec(seed)
	if err != nil {
		return fmt.Errorf("seed data: %w", err)
	}

	return nil
}

func (s *sqliteDB) GetUserLoginDetails(username string) (*LoginDetails, error) {
	if s == nil || s.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	var ld LoginDetails
	row := s.db.QueryRow("SELECT username, authtoken FROM users WHERE username = ?", username)
	if err := row.Scan(&ld.Username, &ld.AuthToken); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %s", username)
		}
		return nil, err
	}
	return &ld, nil
}

func (s *sqliteDB) GetUserCoinDetails(username string) (*CoinDetails, error) {
	if s == nil || s.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	var cd CoinDetails
	row := s.db.QueryRow("SELECT username, coins FROM users WHERE username = ?", username)
	if err := row.Scan(&cd.Username, &cd.Coins); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %s", username)
		}
		return nil, err
	}
	return &cd, nil
}
