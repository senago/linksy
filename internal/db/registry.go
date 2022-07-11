package db

import (
	"github.com/senago/linksy/internal/customtype"
)

type Registry struct {
	ShortenerRepository ShortenerRepository
}

func NewMemoryRegistry() *Registry {
	return nil
}

func NewPostgresRegistry(dbConn *customtype.DBConn) *Registry {
	return &Registry{
		ShortenerRepository: NewShortenerPostgresRepository(dbConn),
	}
}
