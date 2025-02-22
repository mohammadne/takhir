package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadne/zanbil/internal/entities"
	"go.uber.org/zap"
)

func NewLanguage(router fiber.Router, logger *zap.Logger) {
	middleware := &language{
		logger: logger,
	}

	router.Use(middleware.fetchLanguage)
}

type language struct {
	logger *zap.Logger
}

func (l *language) fetchLanguage(c *fiber.Ctx) error {
	languageBytes := c.Request().Header.Peek("language")
	language := entities.ToLanguage(string(languageBytes))
	c.Locals("language", language)
	return c.Next()
}
