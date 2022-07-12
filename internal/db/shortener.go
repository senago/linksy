package db

import (
	"context"
	"sync"

	"github.com/senago/linksy/internal/constants"
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/model/core"
)

const (
	queryCreateURL = `INSERT INTO url ("hash", "value", "created_at", "expires_at") VALUES ($1, $2, $3, $4)
										ON CONFLICT ("hash") DO UPDATE SET
										"hash" = excluded."hash", "value" = excluded."value",
										"created_at" = excluded."created_at", "expires_at" = excluded."expires_at";`
	queryGetURL    = `SELECT "value", "expires_at" FROM url WHERE "hash" = $1;`
	queryDeleteURL = `DELETE FROM url WHERE "hash" = $1;`
)

type ShortenerRepository interface {
	CreateURL(ctx context.Context, url *core.URL) error
	GetURL(ctx context.Context, hash string) (*core.URL, error)
	DeleteURL(ctx context.Context, hash string) error
}

type shortenerMemoryRepository struct {
	store sync.Map
}

func (repo *shortenerMemoryRepository) CreateURL(ctx context.Context, url *core.URL) error {
	repo.store.Store(url.Hash, url)
	return nil
}

func (repo *shortenerMemoryRepository) GetURL(ctx context.Context, hash string) (*core.URL, error) {
	value, ok := repo.store.Load(hash)
	if !ok {
		return nil, constants.ErrDBNotFound
	}
	return value.(*core.URL), nil
}

func (repo *shortenerMemoryRepository) DeleteURL(ctx context.Context, hash string) error {
	repo.store.Delete(hash)
	return nil
}

func NewShortenerMemoryRepository() ShortenerRepository {
	return &shortenerMemoryRepository{}
}

type shortenerPostgresRepository struct {
	dbConn *customtype.DBConn
}

func (repo *shortenerPostgresRepository) CreateURL(ctx context.Context, url *core.URL) error {
	_, err := repo.dbConn.Exec(ctx, queryCreateURL, &url.Hash, &url.Value, &url.CreatedAt, &url.ExpirestAt)
	return err
}

func (repo *shortenerPostgresRepository) GetURL(ctx context.Context, hash string) (*core.URL, error) {
	url := &core.URL{}
	err := repo.dbConn.QueryRow(ctx, queryGetURL, hash).Scan(&url.Value, &url.ExpirestAt)
	return url, wrapErr(err)
}

func (repo *shortenerPostgresRepository) DeleteURL(ctx context.Context, hash string) error {
	_, err := repo.dbConn.Exec(ctx, queryDeleteURL, hash)
	return wrapErr(err)
}

func NewShortenerPostgresRepository(dbConn *customtype.DBConn) ShortenerRepository {
	return &shortenerPostgresRepository{dbConn: dbConn}
}
