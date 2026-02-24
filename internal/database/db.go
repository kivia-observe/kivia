package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseurl string) *pgxpool.Pool {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(databaseurl)

	if err != nil {

		log.Fatal("Unable to parse database pool config", err)
	}

	poolConfig.MaxConns = 100
	poolConfig.MinConns = 10
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MinIdleConns = 5
	poolConfig.MaxConnIdleTime = time.Minute * 10

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)

	if err != nil {

		log.Fatal("Unable to create database pool", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Database ping failed", err)
	}
	return pool
}
