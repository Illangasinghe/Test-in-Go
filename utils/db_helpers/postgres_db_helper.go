package db_helpers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectPostgres connects to the PostgreSQL database using the URL provided in the configuration
func ConnectPostgres(connectionString string) error {
	var err error

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("could not connect to the PostgreSQL database: %v", err)
	}

	// Test the connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("could not ping the PostgreSQL database: %v", err)
	}

	log.Println("Connected to PostgreSQL database successfully.")
	return nil
}

// ClosePostgres closes the PostgreSQL connection
func ClosePostgres() error {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			return fmt.Errorf("could not close the PostgreSQL database: %v", err)
		}
		log.Println("Closed PostgreSQL database connection.")
	}
	return nil
}
