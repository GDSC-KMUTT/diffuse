package fiber

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"backend/types/response"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Case of *fiber.Error.
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(&response.ErrorResponse{
			Success: false,
			Code:    strings.ReplaceAll(strings.ToUpper(e.Error()), " ", "_"),
			Message: e.Error(),
			Error:   e.Error(),
		})
	}

	if e, ok := err.(*response.ErrorInstance); ok {
		if e.Code == "" {
			e.Code = "GENERIC_ERROR"
		}

		if e.Err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
				Success: false,
				Code:    e.Code,
				Message: e.Message,
				Error:   e.Err.Error(),
			})
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Code:    e.Code,
			Message: e.Message,
		})
	}

	// Case of validator.ValidationErrors
	if e, ok := err.(validator.ValidationErrors); ok {
		var lists []string
		for _, err := range e {
			/*
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
			*/

			lists = append(lists, err.Field()+" ("+err.Tag()+")")
		}

		message := strings.Join(lists[:], ", ")

		return ctx.Status(fiber.StatusBadRequest).JSON(&response.ErrorResponse{
			Success: false,
			Code:    "VALIDATION_FAILED",
			Message: "Validation failed on field " + message,
			Error:   e.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(&response.ErrorResponse{
		Success: false,
		Code:    "UNKNOWN_SERVER_SIDE_ERROR",
		Message: "Unknown server side error",
		Error:   err.Error(),
	})
}
