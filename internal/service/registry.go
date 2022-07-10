package service

import (
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/db"
)

type Registry struct{}

func NewRegistry(log *customtype.Logger, repository *db.Repository) *Registry {
	registry := &Registry{}

	return registry
}
