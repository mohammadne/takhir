package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"github.com/mohammadne/zanbil/internal/api/http/i18n"
	"github.com/mohammadne/zanbil/internal/entities"
	"github.com/mohammadne/zanbil/internal/usecases"
)

func NewProducts(router fiber.Router, logger *zap.Logger, i18n i18n.I18N,
	categoriesUsecase usecases.Categories,
) {
	handler := &products{
		logger:            logger,
		i18n:              i18n,
		categoriesUsecase: categoriesUsecase,
	}

	products := router.Group("products")
	products.Get("/", handler.listProducts)
	products.Get("/:id", handler.retrieveProduct)
}

type products struct {
	logger *zap.Logger
	i18n   i18n.I18N

	// usecases
	categoriesUsecase usecases.Categories
}

func (h *products) listProducts(c fiber.Ctx) error {
	listOptions := entities.QueryParamsToListOptions(c.Queries())

	h.logger.Info("queries", zap.Any("a", listOptions))

	products := []entities.Product{
		{ID: 201, Name: "T-shirt", Price: "$19", Description: "Comfortable fabric"},
		{ID: 202, Name: "Headphones1", Price: "$99", Description: "Noise-cancelling"},
		{ID: 203, Name: "Headphones2", Price: "$99", Description: "Noise-cancelling"},
		{ID: 204, Name: "Headphones3", Price: "$99", Description: "Noise-cancelling"},
		{ID: 205, Name: "Headphones4", Price: "$99", Description: "Noise-cancelling"},
		{ID: 206, Name: "Headphones5", Price: "$99", Description: "Noise-cancelling"},
		{ID: 207, Name: "Headphones6", Price: "$99", Description: "Noise-cancelling"},
		{ID: 208, Name: "Headphones7", Price: "$99", Description: "Noise-cancelling"},
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"products":    products,
		"total_count": len(products),
	})
}

func (h *products) retrieveProduct(c fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}
