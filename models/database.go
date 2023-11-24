package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type RepoDB struct {
	*sql.DB
}

func ConnectDB(dataSourceName string) (*RepoDB, error) {
	fmt.Print(dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to database!")
	return &RepoDB{db}, nil
}

func Close(db *RepoDB) error {
	err := db.Close()
	if err != nil {
		log.Fatalf("failed to close database connection: %v", err)
	}
	fmt.Println("Database connection closed!")
	return err
}
