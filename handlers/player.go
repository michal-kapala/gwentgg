package handlers

import (
	"fmt"
	"gwentgg/components"
	"gwentgg/config"
	"gwentgg/db"
	"gwentgg/db/models"
	"gwentgg/services/cards"

	"github.com/gofiber/fiber/v3"
)

func PlayerHandler(c fiber.Ctx) error {
	userID := c.Params("user")
	token := c.Cookies("access_token")
	expired := c.Cookies("token_expired", "true")
	if token == "" || expired == "true" {
		return c.Redirect().To("/login")
	}
	database := db.Get(c)

	var card models.CardDefinition
	database.First(&card)
	if card.ID == "" {
		fmt.Println("Fetching card definitions...")
		cfg := config.Get(c)
		resp, err := cards.Get(cfg, token)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Card definition request error, see server log for details.")
		}
		if resp.Status() != "200 OK" {
			return c.Status(fiber.StatusBadRequest).SendString("Card definition request failed, check your credentials.")
		}
		defs := resp.Result().(*cards.CardDefList)
		cardDefs, err := cards.ToModels(defs.Items)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Card data transformation error, see server log for details.")
		}
		database.CreateInBatches(&cardDefs, 250)
	}

	var user models.User
	database.First(&user, "id = ?", userID)
	if user.ID == "" {
		return Render(c, components.NotFound())
	}
	return Render(c, components.PlayerProfile(&user))
}
