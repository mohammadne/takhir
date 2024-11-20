package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Items interface {
	List(c *fiber.Ctx) error
}

func NewItems() Items {
	return &items{}
}

type items struct{}

func (i *items) List(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
