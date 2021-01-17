package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var (
	// DBCon will be an instance of the database connection
	DBCon *pgx.Conn
)

// GetConnection will create the connection
func GetConnection() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	conConfig, err := pgx.ParseConfig(os.Getenv("DB_URL"))

	if err != nil {
		panic(err)
	}

	DBCon, err = pgx.ConnectConfig(context.Background(), conConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
