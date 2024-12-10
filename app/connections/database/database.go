package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	cfg "github.com/nannigalaxy/go-rest-api/app/config"
)

func connect() (*pgxpool.Pool, context.Context) {
	// Connection URL
	dbURL := cfg.Config.DBUrl

	// Configure connection pool
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Panicf("unable to parse database URL: %v", err)
		return nil, nil
	}

	// Set connection pool settings
	config.MaxConns = 20                      // Maximum number of connections
	config.MinConns = 5                       // Minimum number of connections
	config.MaxConnLifetime = time.Hour        // Max lifetime of a connection
	config.MaxConnIdleTime = 30 * time.Minute // Max idle time of a connection

	// Create the pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Panicf("failed to create connection pool: %v", err)
		return nil, nil
	}

	// Test the connection
	err = pool.Ping(context.Background())
	if err != nil {
		pool.Close()
		log.Panicf("failed to connect to the database: %v", err)
		return nil, nil
	}

	ctx := context.Background()

	_, err2 := pool.Exec(ctx, "DISCARD ALL")
	if err2 != nil {
		log.Printf("Error resetting connection state: %v", err)
	}

	log.Println("Successfully connected to the database")
	return pool, ctx
}

var DBConnection, DBContext = connect()
