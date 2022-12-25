package db

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Database struct {
	Client *sqlx.DB
}

// Singleton object - where
// File descriptor, total limit on
// New ports will also open
var (
	dbClient *Database
)

func NewDatabase() (*Database, error) {
	fmt.Println(dbClient)
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
		// Whenever we are accessing a pointer ,  check null pointer exception
		dbClient = &Database{
			Client: conn,
		}
	}
	return dbClient, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}
