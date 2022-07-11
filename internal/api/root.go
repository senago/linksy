package api

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/senago/linksy/internal/api/controller"
	"github.com/senago/linksy/internal/constants"
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/db"
)

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	} else if e, ok := err.(*constants.CodedError); ok {
		code = e.Code
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(code).SendString(err.Error())
}

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
			ErrorHandler: errorHandler,
			JSONEncoder:  sonic.Marshal,
			JSONDecoder:  sonic.Unmarshal,
		}),
	}

	registry, err := controller.NewRegistry(log, dbRegistry)
	if err != nil {
		return nil, err
	}

	api := svc.router.Group("/api", recover.New())

	api.Post("/shorten", registry.ShortenerController.Shorten)
	api.Get("/retrieve", registry.ShortenerController.Retrieve)

	return svc, nil
}
