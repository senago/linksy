package db

import (
	"context"

	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/model/core"
)

const (
	queryCreateURL = `INSERT INTO url ("hash", "value", "created_at", "expires_at") VALUES ($1, $2, $3, $4);`
	queryGetURL    = `SELECT "value", "expires_at" FROM url WHERE "hash" = $1;`
	queryDeleteURL = `DELETE FROM url WHERE "hash" = $1;`
)

type ShortenerRepository interface {
	CreateURL(ctx context.Context, url *core.URL) error
	GetURL(ctx context.Context, hash string) (*core.URL, error)
	DeleteURL(ctx context.Context, hash string) error
}

type urlRepositoryImpl struct {
	dbConn *customtype.DBConn
}

func (repo *urlRepositoryImpl) CreateURL(ctx context.Context, url *core.URL) error {
	_, err := repo.dbConn.Exec(ctx, queryCreateURL, &url.Hash, &url.Value, &url.CreatedAt, &url.ExpirestAt)
	return err
}

func (repo *urlRepositoryImpl) GetURL(ctx context.Context, hash string) (*core.URL, error) {
	url := &core.URL{}
	err := repo.dbConn.QueryRow(ctx, queryGetURL, hash).Scan(&url.Value, &url.ExpirestAt)
	return url, wrapErr(err)
}

func (repo *urlRepositoryImpl) DeleteURL(ctx context.Context, hash string) error {
	_, err := repo.dbConn.Exec(ctx, queryDeleteURL, hash)
	return wrapErr(err)
}

func NewShortenerPostgresRepository(dbConn *customtype.DBConn) ShortenerRepository {
	return &urlRepositoryImpl{dbConn: dbConn}
}
