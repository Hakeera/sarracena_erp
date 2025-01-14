// Package database handles the connection pool management for PostgreSQL.
package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

// InitDatabase initializes the PostgreSQL connection pool using the connection string
// provided in the DATABASE_URL environment variable. It logs a fatal error if the connection cannot be established.
func InitDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	var err error

	// Log the DSN (masked for security purposes)
	log.Printf("Initializing database connection pool with DSN: %s\n", maskDSN(dsn))

	// Create the database connection pool with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbPool, err = pgxpool.New(ctx, dsn)
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
	log.Println("Database connection pool retrieved successfully")
	return dbPool
}

// CloseDatabase closes the PostgreSQL connection pool, releasing all resources.
func CloseDatabase() {
	if dbPool != nil {
		dbPool.Close()
		fmt.Println("Database connection pool closed.")
	}
}

// maskDSN masks sensitive parts of the DSN (Database Connection String) for logging.
// This function replaces sensitive information like username and password with "****".
func maskDSN(dsn string) string {
	if dsn == "" {
		return "DSN is empty"
	}
	// Simple example of masking the DSN. You can enhance this as needed.
	// Assuming the DSN format is "postgres://username:password@host:port/dbname"
	// We will mask the username and password here.
	dsn = maskSubstring(dsn, "postgres://", "")
	dsn = maskSubstring(dsn, ":", "****:")
	return dsn
}

// maskSubstring replaces part of the DSN string with the provided mask.
func maskSubstring(dsn, old, mask string) string {
	if dsn == "" || old == "" {
		return dsn
	}
	if idx := len(dsn) - len(old); idx >= 0 && dsn[idx:] == old {
		return dsn[:idx] + mask
	}
	return dsn
}
