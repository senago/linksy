package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/model/dto"
	"github.com/senago/linksy/internal/service"
)

type ShortenerController struct {
	log      *customtype.Logger
	registry *service.Registry
}

func (c *ShortenerController) Shorten(ctx *fiber.Ctx) error {
	request := &dto.ShortenRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.JSON(nil)
}

func (c *ShortenerController) Retrieve(ctx *fiber.Ctx) error {
	request := &dto.RetrieveRequest{}
	if err := ctx.QueryParser(request); err != nil {
		return err
	}

	return ctx.JSON(nil)
}

func NewShortenerController(log *customtype.Logger, registry *service.Registry) *ShortenerController {
	return &ShortenerController{log: log, registry: registry}
}
