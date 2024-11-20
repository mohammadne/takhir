package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Healthz interface {
	Liveness(c *fiber.Ctx) error
	Readiness(c *fiber.Ctx) error
}

func NewHealthz() Healthz {
	return &healthz{}
}

type healthz struct{}

func (h *healthz) Liveness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *healthz) Readiness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
