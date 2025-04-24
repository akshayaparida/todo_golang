package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/akshayaparida/todo_golang/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dbUser := config.GetEnv("DB_USER")
	dbPass := config.GetEnv("DB_PASSWORD")
	dbName := config.GetEnv("DB_NAME")
	dbHost := config.GetEnv("DB_HOST")
	dbPort := config.GetEnv("DB_PORT")
	sslMode := config.GetEnv("SSL_MODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslMode)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}

	log.Println(" Connected to PostgreSQL!")
	DB = pool
}
