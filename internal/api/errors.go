package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senago/linksy/internal/constants"
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
