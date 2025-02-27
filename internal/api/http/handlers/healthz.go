package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func NewHealthz(router fiber.Router, logger *zap.Logger) {
	handler := &healthz{logger: logger}

	healthz := router.Group("healthz")
	healthz.Get("/liveness", handler.liveness)
	healthz.Get("/readiness", handler.readiness)
}

type healthz struct {
	logger *zap.Logger
}

func (h *healthz) liveness(c fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *healthz) readiness(c fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
