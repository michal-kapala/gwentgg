package handlers

import (
	"fmt"
	"gwentgg/components"
	"gwentgg/config"
	"gwentgg/db"
	"gwentgg/db/models"
	"gwentgg/services/cards"
	"gwentgg/services/games"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PlayerHandler(c fiber.Ctx) error {
	userID := c.Params("user")
	token := c.Cookies("access_token")
	expired := c.Cookies("token_expired", "true")
	if token == "" || expired == "true" {
		return c.Redirect().To("/login")
	}
	database := db.Get(c)
	cfg := config.Get(c)

	var card models.CardDefinition
	database.First(&card)
	if card.ID == "" {
		fmt.Println("Fetching card definitions...")
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

	resp, err := games.GetPage(cfg, token, 1)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Latest games request error, see server log for details.")
	}

	if resp.Status() != "200 OK" {
		return c.Status(fiber.StatusBadRequest).SendString("Latest games request failed, check your credentials.")
	}

	latestGames := resp.Result().(*games.GameList)
	games := []models.Game{}

	for _, game := range latestGames.Items {
		model, err := game.ToModel()
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Card data transformation error, see server log for details.")
		}
		games = append(games, model)
	}

	database.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Clauses(clause.OnConflict{DoNothing: true}).
		Save(&games)

	var user models.User
	database.First(&user, "id = ?", userID)
	if user.ID == "" {
		return Render(c, components.NotFound())
	}
	return Render(c, components.PlayerProfile(&user, games))
}
