package shortener

import (
	"context"

	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/model/core"
)

const (
	queryCreateForum = `INSERT INTO url ("hash", "value", "created_at", "expires_at") VALUES ($1, $2, $3, $4);`
)

type urlRepositoryImpl struct {
	dbConn *customtype.DBConn
}

func (repo *urlRepositoryImpl) CreateURL(ctx context.Context, url *core.URL) error {
	_, err := repo.dbConn.Exec(ctx, queryCreateForum, &url.Hash, &url.Value, &url.CreatedAt, &url.ExpirestAt)
	return err
}

func (repo *urlRepositoryImpl) GetURL(ctx context.Context, hash string) (*core.URL, error) {
	return nil, nil
}

func NewURLPostgresRepository(dbConn *customtype.DBConn) Repository {
	return &urlRepositoryImpl{dbConn: dbConn}
}
