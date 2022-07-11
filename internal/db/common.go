package db

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/senago/linksy/internal/constants"
)

func wrapErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return constants.ErrDBNotFound
	}
	return err
}
