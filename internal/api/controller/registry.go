package controller

import (
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/db"
	"github.com/senago/linksy/internal/service"
)

type Registry struct {
	ShortenerController *ShortenerController
}

func NewRegistry(log *customtype.Logger, dbRegistry *db.Registry) (*Registry, error) {
	serviceRegistry := service.NewRegistry(log, dbRegistry)

	return &Registry{
		ShortenerController: NewShortenerController(log, serviceRegistry),
	}, nil
}
