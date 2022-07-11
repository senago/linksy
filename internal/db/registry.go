package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/senago/linksy/internal/db/shortener"
	"github.com/spf13/viper"
)

const (
	dbTypeMemory   = "memory"
	dbTypePostgres = "postgres"
)

type Registry struct {
	ShortenerRepository shortener.Repository
}

func NewRegistry(dbType string) (*Registry, error) {
	repository := &Registry{}

	switch dbType {
	case dbTypeMemory:
	case dbTypePostgres:
		dbPool, err := pgxpool.Connect(context.Background(), viper.GetString("postgres.connection_string"))
		if err != nil {
			return nil, fmt.Errorf("failed to connect to the postgres database: %s", err)
		}
		// defer dbPool.Close()

		repository.ShortenerRepository = shortener.NewURLPostgresRepository(dbPool)

	default:
		return nil, fmt.Errorf("unexpected db type: [%s]", dbType)
	}

	return repository, nil
}