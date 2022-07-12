package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/senago/linksy/internal/customtype"
	"github.com/senago/linksy/internal/model/dto"
	"github.com/senago/linksy/internal/service"
)

type ShortenerController struct {
	log *customtype.Logger
	svc *service.Registry
}

func (c *ShortenerController) Shorten(ctx *fiber.Ctx) error {
	request := &dto.ShortenRequest{}
	if err := Bind(ctx, request, ctx.BodyParser); err != nil {
		return err
	}

	response, err := c.svc.ShortenerService.Shorten(context.Background(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *ShortenerController) Retrieve(ctx *fiber.Ctx) error {
	request := &dto.RetrieveRequest{}
	if err := Bind(ctx, request, ctx.QueryParser); err != nil {
		return err
	}

	response, err := c.svc.ShortenerService.Retrieve(context.Background(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func NewShortenerController(log *customtype.Logger, svc *service.Registry) *ShortenerController {
	return &ShortenerController{log: log, svc: svc}
}
