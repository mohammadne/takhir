package handlers

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

//go:embed templates/*
var files embed.FS

func NewTemplates(router fiber.Router, logger *zap.Logger) {
	handler := &home{
		logger: logger,
	}

	homeBytes, err := files.ReadFile("templates/home.html")
	if err != nil {
		logger.Panic("", zap.Error(err))
	}
	handler.home = string(homeBytes)

	router.Get("/", handler.listProducts)
}

type home struct {
	logger *zap.Logger

	// templates
	home string
}

func (h *home) listProducts(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html")
	return ctx.Status(http.StatusOK).SendString(h.home)
}
