package handlers

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"go.uber.org/zap"

	"github.com/mohammadne/zanbil/internal/entities"
)

//go:embed static/*
var staticFS embed.FS

//go:embed templates/*
var templatesFS embed.FS

// PageData represents the data passed to the template
type PageData struct {
	Languages  []Language
	Categories []entities.Category
}

type Language struct {
	Name string
	Code string
}

func NewTemplates(router fiber.Router, logger *zap.Logger) {
	handler := &templates{
		logger: logger,
	}

	router.Get("/*", static.New("", static.Config{
		FS:     staticFS,
		Browse: true,
	}))

	handler.homeTpl = template.Must(template.ParseFS(templatesFS, "templates/home.html"))
	router.Get("/home", handler.home)

	handler.productTpl = template.Must(template.ParseFS(templatesFS, "templates/product.html"))
	router.Get("/product/:id", handler.product)
}

type templates struct {
	logger *zap.Logger

	// templates
	homeTpl    *template.Template
	productTpl *template.Template
}

func (h *templates) home(ctx fiber.Ctx) error {
	// Dummy categories
	categories := []entities.Category{
		{ID: 1, Name: "Electronics"},
		{ID: 2, Name: "Clothing"},
		{ID: 3, Name: "Books"},
	}

	// Render the template into a string
	var renderedHTML bytes.Buffer
	err := h.homeTpl.Execute(&renderedHTML, PageData{
		Languages: []Language{
			{Name: "English", Code: "en"},
			{Name: "Persian", Code: "fa"},
		},
		Categories: categories,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Template rendering error")
	}

	ctx.Set("Content-Type", "text/html")
	return ctx.Status(http.StatusOK).SendString(renderedHTML.String())
}

func (h *templates) product(ctx fiber.Ctx) error {
	productID, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil || productID <= 0 {
		return ctx.Status(http.StatusBadRequest).SendString("Invalid id")
	}

	// Render the template into a string
	var renderedHTML bytes.Buffer
	err = h.productTpl.Execute(&renderedHTML, entities.Product{
		ID:          productID,
		Name:        "Gaming Mouse",
		Price:       "$49",
		Description: "High precision sensor with RGB lighting.",
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Template rendering error")
	}

	ctx.Set("Content-Type", "text/html")
	return ctx.Status(http.StatusOK).SendString(renderedHTML.String())
}
