package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/mohammadne/zanbil/internal/api/http/i18n"
	"github.com/mohammadne/zanbil/internal/api/http/models"
	"github.com/mohammadne/zanbil/internal/usecases"
)

func NewCategories(router fiber.Router, logger *zap.Logger, i18n i18n.I18N,
	categoriesUsecase usecases.Categories,
) {
	handler := &categories{
		logger:            logger,
		i18n:              i18n,
		categoriesUsecase: categoriesUsecase,
	}

	categories := router.Group("categories")
	categories.Get("/", handler.list)
}

type categories struct {
	logger *zap.Logger
	i18n   i18n.I18N

	// usecases
	categoriesUsecase usecases.Categories
}

func (c *categories) list(ctx *fiber.Ctx) error {
	response := &models.Response{}

	categories, failure := c.categoriesUsecase.AllCategories(ctx.Context())
	if failure != nil {
		response.Message = c.i18n.Translate("categories.list.error", "fa")
		return response.Write(ctx, fiber.StatusInternalServerError)
	}

	response.Data = map[string]any{"categories": categories}
	response.Message = c.i18n.Translate("categories.list.success", "fa")
	return response.Write(ctx, fiber.StatusOK)
}
