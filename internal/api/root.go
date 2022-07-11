package api

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/senago/linksy/internal/api/controller"
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/db"
)

type APIService struct {
	log    *customtype.Logger
	router *fiber.App
}

func (svc *APIService) Serve(addr string) {
	svc.log.Fatal(svc.router.Listen(addr))
}

func (svc *APIService) Shutdown(ctx context.Context) error {
	return svc.router.Shutdown()
}

func NewAPIService(log *customtype.Logger, dbRegistry *db.Registry) (*APIService, error) {
	svc := &APIService{
		log: log,
		router: fiber.New(fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
		}),
	}

	registry, err := controller.NewRegistry(log, dbRegistry)
	if err != nil {
		return nil, err
	}

	api := svc.router.Group("/api")

	api.Post("/shorten", registry.ShortenerController.Shorten)
	api.Get("/retrieve", registry.ShortenerController.Retrieve)

	return svc, nil
}
