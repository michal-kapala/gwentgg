package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func IndexHandler(c fiber.Ctx) error {
	return c.Redirect().To("/login")
}
