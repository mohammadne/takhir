package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadne/zanbil/internal/api/http/i18n"
	"github.com/mohammadne/zanbil/internal/entities"
	"github.com/mohammadne/zanbil/internal/usecases"
	"go.uber.org/zap"
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

func (h *products) listProducts(c *fiber.Ctx) error {
	products := []entities.Product{
		{ID: 201, Name: "T-shirt", Price: "$19", Description: "Comfortable fabric"},
		{ID: 202, Name: "Headphones", Price: "$99", Description: "Noise-cancelling"},
	}

	return c.Status(http.StatusOK).JSON(products)
}

func (h *products) retrieveProduct(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNotImplemented)
}
