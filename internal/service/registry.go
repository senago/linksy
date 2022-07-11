package service

import (
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/db"
)

type Registry struct {
	ShortenerService ShortenerService
}

func NewRegistry(log *customtype.Logger, dbRegistry *db.Registry) *Registry {
	return &Registry{
		ShortenerService: NewShortenerService(log, dbRegistry),
	}
}
