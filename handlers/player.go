package handlers

import (
	"github.com/gofiber/fiber/v3"
	"gwentgg/components"
	"gwentgg/db"
	models "gwentgg/db/models"
)

func PlayerHandler(c fiber.Ctx) error {
	userID := c.Params("user")
	token := c.Cookies("access_token")
	expired := c.Cookies("token_expired", "true")
	if token == "" || expired == "true" {
		return c.Redirect().To("/login")
	}
	database := db.Get(c)
	var user models.User
	database.First(&user, "id = ?", userID)
	if user.ID == "" {
		return Render(c, components.NotFound())
	}
	return Render(c, components.PlayerProfile(&user))
}
