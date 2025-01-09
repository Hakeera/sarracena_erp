// Package database handles the connection pool management for PostgreSQL.
package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

// InitDatabase initializes the PostgreSQL connection pool using the connection string
// provided in the `DATABASE_URL` environment variable. It logs a fatal error if the connection cannot be established.
func InitDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	var err error

	dbPool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	fmt.Println("Database connection pool created successfully!")
}

// GetDB retrieves the current PostgreSQL connection pool instance.
// Logs a fatal error if the connection is not initialized.
func GetDB() *pgxpool.Pool {
	if dbPool == nil {
		log.Fatal("Database connection is not initialized")
	}
	return dbPool
}

// CloseDatabase closes the PostgreSQL connection pool, releasing all resources.
func CloseDatabase() {
	if dbPool != nil {
		dbPool.Close()
		fmt.Println("Database connection pool closed.")
	}
}
