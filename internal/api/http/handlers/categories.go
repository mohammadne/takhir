package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/mohammadne/zanbil/internal/api/http/i18n"
	"github.com/mohammadne/zanbil/internal/api/http/models"
	"github.com/mohammadne/zanbil/internal/entities"
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
	categories.Get("/", handler.allCategories)
	categories.Get("/:id/products", handler.listCategoryProducts)
}

type categories struct {
	logger *zap.Logger
	i18n   i18n.I18N

	// usecases
	categoriesUsecase usecases.Categories
}

func (c *categories) allCategories(ctx *fiber.Ctx) error {
	response := &models.Response{}
	language, _ := ctx.Locals("language").(entities.Language)

	categories, error := c.categoriesUsecase.AllCategories(ctx.Context())
	if error != nil {
		response.Message = c.i18n.Translate("categories.all_categories.error", language)
		return response.Write(ctx, fiber.StatusInternalServerError)
	}

	response.Data = map[string]any{"categories": categories}
	response.Message = c.i18n.Translate("categories.all_categories.success", language)
	return response.Write(ctx, fiber.StatusOK)
}

func (c *categories) listCategoryProducts(ctx *fiber.Ctx) error {
	response := &models.Response{}
	language, _ := ctx.Locals("language").(entities.Language)

	categoryID, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil || categoryID <= 0 {
		response.Message = c.i18n.Translate("categories.list_category_products.invalid_category_id", language)
		return response.Write(ctx, fiber.StatusBadRequest)
	}

	categories, error := c.categoriesUsecase.ListCategoryProducts(ctx.Context(), categoryID, nil)
	if error != nil {
		response.Message = c.i18n.Translate("categories.list_category_products.error", language)
		return response.Write(ctx, fiber.StatusInternalServerError)
	}

	response.Data = map[string]any{"categories": categories}
	response.Message = c.i18n.Translate("categories.list_category_products.success", language)
	return response.Write(ctx, fiber.StatusOK)
}
