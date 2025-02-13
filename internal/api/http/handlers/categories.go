package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewCategories(router fiber.Router, logger *zap.Logger) {
	handler := &categories{logger: logger}

	categories := router.Group("categories")
	categories.Get("/", handler.list)
}

type categories struct {
	logger *zap.Logger
}

func (i *categories) list(c *fiber.Ctx) error {
	return c.SendString("categories")
}
