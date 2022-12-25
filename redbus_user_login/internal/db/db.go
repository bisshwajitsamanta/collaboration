package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Database struct {
	Client *sqlx.DB
}

/*
									==== SingleTon Implementation ===
 Singleton object creation, so that if one connection dbClient presents no need to create other db connections.
 It saves us from having out of File descriptor issue and also ports running out because of say multiple client
 connections to our RDS Postgresql Database.

*/
var (
	dbClient *Database
)

func NewDatabase() (*Database, error) {
	if dbClient == nil {
		connectionString := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_TABLE"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("SSL_MODE"),
		)
		conn, err := sqlx.Connect("postgres", connectionString)
		if err != nil {
			return &Database{}, fmt.Errorf("Could not connect to the Database:%w", err)
		}
		dbClient = &Database{
			Client: conn,
		}
	}
	return dbClient, nil
}
