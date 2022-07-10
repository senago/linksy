package controller

import (
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/db"
	"github.com/senago/linksy/internal/service"
)

type Registry struct {
	ShortenerController *ShortenerController
}

func NewRegistry(log *customtype.Logger, dbConn *customtype.DBConn) (*Registry, error) {
	repository, err := db.NewRepository(dbConn)
	if err != nil {
		return nil, err
	}

	serviceRegistry := service.NewRegistry(log, repository)

	return &Registry{
		ShortenerController: NewShortenerController(log, serviceRegistry),
	}, nil
}
