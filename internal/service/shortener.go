package service

import (
	"context"
	"time"

	"github.com/senago/linksy/internal/constants"
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/db"
	"github.com/senago/linksy/internal/model/core"
	"github.com/senago/linksy/internal/model/dto"
	"github.com/senago/linksy/internal/util"
)

const (
	LinkExpirationTime = time.Hour * 24 * 365
)

type ShortenerService interface {
	Shorten(ctx context.Context, request *dto.ShortenRequest) (*dto.ShortenResponse, error)
	Retrieve(ctx context.Context, request *dto.RetrieveRequest) (*dto.RetrieveResponse, error)
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
		ExpirestAt: timeNow.Add(LinkExpirationTime),
	}

	if err := svc.db.ShortenerRepository.CreateURL(ctx, url); err != nil {
		return nil, err
	}

	return &dto.ShortenResponse{Hash: url.Hash}, nil
}

func (svc *shortenerServiceImpl) Retrieve(ctx context.Context, request *dto.RetrieveRequest) (*dto.RetrieveResponse, error) {
	url, err := svc.db.ShortenerRepository.GetURL(ctx, request.Hash)
	if err != nil {
		return nil, err
	}

	if url.ExpirestAt.Before(time.Now()) {
		if err := svc.db.ShortenerRepository.DeleteURL(ctx, request.Hash); err != nil {
			if err != constants.ErrDBNotFound {
				return nil, err
			}
		}
		return nil, constants.ErrLinkExpired
	}

	return &dto.RetrieveResponse{URL: url.Value}, nil
}

func NewShortenerService(log *customtype.Logger, db *db.Registry) ShortenerService {
	return &shortenerServiceImpl{log: log, db: db}
}
