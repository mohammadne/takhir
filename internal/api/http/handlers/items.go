package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewItems(router fiber.Router, logger *zap.Logger) {
	handler := &items{logger: logger}

	items := router.Group("items")
	items.Get("/", handler.list)
	// items.Get("/:id", server.getItem)
}

type items struct {
	logger *zap.Logger
}

func (i *items) list(c *fiber.Ctx) error {
	return c.SendString("items")
}
