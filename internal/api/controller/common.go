package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/senago/linksy/internal/constants"
)

type Parser func(out interface{}) error

var validate = validator.New()

func ValidateStruct(i interface{}) error {
	if err := validate.Struct(i); err != nil {
		builder := strings.Builder{}
		for _, ve := range err.(validator.ValidationErrors) {
			builder.WriteString(fmt.Sprintf("[%s] didn't satisfy [%s]", ve.Param(), ve.Tag()))
		}
		return constants.NewCodedError(builder.String(), http.StatusBadRequest)
	}
	return nil
}

func Bind(ctx *fiber.Ctx, out interface{}, parsers ...Parser) error {
	for _, parser := range parsers {
		parser(out)
	}

	if err := ValidateStruct(out); err != nil {
		return err
	}

	return nil
}
