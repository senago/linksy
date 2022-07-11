package shortener

import (
	"context"

	"github.com/senago/linksy/internal/model/core"
)

type Repository interface {
	CreateURL(ctx context.Context, url *core.URL) error
	GetURL(ctx context.Context, hash string) (*core.URL, error)
}
