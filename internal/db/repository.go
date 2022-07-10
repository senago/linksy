package db

import (
	"github.com/senago/linksy/internal/customtype"
)

type Repository struct{}

func NewRepository(dbConn *customtype.DBConn) (*Repository, error) {
	repository := &Repository{}

	return repository, nil
}
