package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

// InitDatabase initializes the connection pool to the PostgreSQL database
func InitDatabase() {
	dsn := os.Getenv("DATABASE_URL") // A URL do banco deve estar configurada no .env
	var err error

	dbPool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	fmt.Println("Database connection pool created successfully!")
}

// GetDB returns the current database pool instance
func GetDB() *pgxpool.Pool {
	if dbPool == nil {
		log.Fatal("Database connection is not initialized")
	}
	return dbPool
}

// CloseDatabase closes the connection pool
func CloseDatabase() {
	if dbPool != nil {
		dbPool.Close()
		fmt.Println("Database connection pool closed.")
	}
}
