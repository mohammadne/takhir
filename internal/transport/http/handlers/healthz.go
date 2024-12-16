package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewHealthz(router fiber.Router, logger *zap.Logger) {
	healthz := &healthz{logger: logger}

	router.Get("/liveness", healthz.liveness)
	router.Get("/readiness", healthz.readiness)
}

type healthz struct {
	logger *zap.Logger
}

func (h *healthz) liveness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *healthz) readiness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
