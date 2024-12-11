package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewItems(router fiber.Router, logger *zap.Logger) {
	items := &items{logger: logger}

	group := router.Group("items")
	group.Get("/", items.list)
	// group.Get("/:id", server.getItem)
}

type items struct {
	logger *zap.Logger
}

func (i *items) list(c *fiber.Ctx) error {
	return c.SendString("items")
}
