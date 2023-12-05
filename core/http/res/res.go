package res

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ValidationError(ctx *fiber.Ctx, errors map[string]any) error {
	return ctx.Status(412).JSON(fiber.Map{
		"data":   nil,
		"error":  errors,
		"status": "VALIDATION_ERROR",
	})
}

func Success(ctx *fiber.Ctx, data any) error {
	return ctx.Status(200).JSON(fiber.Map{
		"data":   data,
		"error":  nil,
		"status": "SUCCESS",
	})
}

func BadRequestError(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(400).JSON(fiber.Map{
		"data":   nil,
		"error":  msg,
		"status": "BAD_REQUEST",
	})
}

func Error(ctx *fiber.Ctx, msg string, statusCode int) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"data":  nil,
		"error": msg,
		"status": func() string {
			switch statusCode {
			case http.StatusBadRequest:
				return "BAD_REQUEST"
			case http.StatusPreconditionFailed:
				return "VALIDATION_ERROR"
			default:
				return "ERROR"
			}
		}(),
	})
}
