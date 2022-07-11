package service

import (
	"context"
	"time"

	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/db"
	"github.com/senago/linksy/internal/model/core"
	"github.com/senago/linksy/internal/model/dto"
	"github.com/senago/linksy/internal/util"
)

type ShortenerService interface {
	Shorten(ctx context.Context, request *dto.ShortenRequest) (*dto.ShortenResponse, error)
}

type shortenerServiceImpl struct {
	log *customtype.Logger
	db  *db.Registry
}

func (svc *shortenerServiceImpl) Shorten(ctx context.Context, request *dto.ShortenRequest) (*dto.ShortenResponse, error) {
	timeNow := time.Now()
	url := &core.URL{
		Hash:       util.Shorten(request.URL),
		Value:      request.URL,
		CreatedAt:  timeNow,
		ExpirestAt: timeNow.Add(time.Hour * 72),
	}

	if err := svc.db.ShortenerRepository.CreateURL(ctx, url); err != nil {
		return nil, err
	}

	return &dto.ShortenResponse{Hash: url.Hash}, nil
}

func NewShortenerService(log *customtype.Logger, db *db.Registry) ShortenerService {
	return &shortenerServiceImpl{log: log, db: db}
}
