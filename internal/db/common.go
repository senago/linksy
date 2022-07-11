package db

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/senago/linksy/internal/constants"
)

var (
	mapping = map[error]error{pgx.ErrNoRows: constants.ErrDBNotFound}
)

func wrapErr(err error) error {
	for k, v := range mapping {
		if errors.Is(err, k) {
			return v
		}
	}
	return err
}
