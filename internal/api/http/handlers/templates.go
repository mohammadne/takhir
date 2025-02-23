package handlers

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadne/zanbil/internal/entities"
	"go.uber.org/zap"
)

//go:embed templates/*
var templatesFS embed.FS

// PageData represents the data passed to the template
type PageData struct {
	Categories []entities.Category
}

func NewTemplates(router fiber.Router, logger *zap.Logger) {
	handler := &templates{
		logger: logger,
	}

	handler.homeTpl = template.Must(template.ParseFS(templatesFS, "templates/home.html"))
	handler.productTpl = template.Must(template.ParseFS(templatesFS, "templates/product.html"))

	router.Get("/home", handler.home)
	router.Get("/product/:id", handler.product)
}

type templates struct {
	logger *zap.Logger

	// templates
	homeTpl    *template.Template
	productTpl *template.Template
}

func (h *templates) home(ctx *fiber.Ctx) error {
	// Dummy categories
	categories := []entities.Category{
		{ID: 1, Name: "Electronics"},
		{ID: 2, Name: "Clothing"},
		{ID: 3, Name: "Books"},
	}

	// Render the template into a string
	var renderedHTML bytes.Buffer
	err := h.homeTpl.Execute(&renderedHTML, PageData{Categories: categories})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Template rendering error")
	}

	ctx.Set("Content-Type", "text/html")
	return ctx.Status(http.StatusOK).SendString(renderedHTML.String())
}

func (h *templates) product(ctx *fiber.Ctx) error {
	productID, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil || productID <= 0 {
		return ctx.Status(http.StatusBadRequest).SendString("Invalid id")
	}

	// Fetch product from DB or predefined data
	product := entities.Product{
		ID:          productID,
		Name:        "Gaming Mouse",
		Price:       "$49",
		Description: "High precision sensor with RGB lighting.",
	}

	// Render the template into a string
	var renderedHTML bytes.Buffer
	err = h.productTpl.Execute(&renderedHTML, product)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Template rendering error")
	}

	ctx.Set("Content-Type", "text/html")
	return ctx.Status(http.StatusOK).SendString(renderedHTML.String())
}
