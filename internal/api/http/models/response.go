package models

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func (response *Response) Write(ctx *fiber.Ctx, statusCode int) error {
	ctx.Set("Content-Type", "application/json")
	return ctx.Status(statusCode).JSON(&response)
}
