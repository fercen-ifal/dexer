package models

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDatabase() (*pgxpool.Pool, error) {
	var (
		pool *pgxpool.Pool
		err  error
	)

	connect := func() error {
		pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		return err
	}

	backoffStrategy := backoff.NewExponentialBackOff()
	backoffStrategy.MaxElapsedTime = time.Minute

	err = backoff.Retry(connect, backoffStrategy)
	if err != nil {
		log.Printf("Não foi possível se conectar ao banco de dados: %e", err)
		return nil, err
	}

	// TODO: Remover teste SQL

	_, err = pool.Query(context.Background(), "CREATE TABLE IF NOT EXISTS test (value STRING PRIMARY KEY)")
	if err != nil {
		log.Printf("Não foi possível criar tabela teste: %e", err)
		return nil, err
	}

	return pool, nil
}
