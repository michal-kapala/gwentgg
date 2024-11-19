package handlers

import (
	"fmt"
	"gwentgg/components/pages"
	"gwentgg/config"
	"gwentgg/db"
	"gwentgg/db/models"
	"gwentgg/services/cards"
	"gwentgg/services/games"
	"gwentgg/services/stats"
	"math"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PlayerHandler(c fiber.Ctx) error {
	userID := c.Params("user")
	token := c.Cookies("access_token")
	seasonID := c.Cookies("current_season", "")
	expired := c.Cookies("token_expired", "true")
	if token == "" || expired == "true" {
		return c.Redirect().To("/login")
	}
	database := db.Get(c)
	cfg := config.Get(c)

	var card models.CardDefinition
	database.First(&card)
	// new season - BC patch available
	stale := time.Now().Month() != card.UpdatedAt.Month()
	if card.ID == "" || stale {
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
		chunkSize := 250
		chunks := int(math.Ceil(float64(len(cardDefs)) / float64(chunkSize)))
		if stale {
			var chunk []models.CardDefinition
			for idx := 0; idx < chunks; idx++ {
				if idx == chunks-1 {
					chunk = cardDefs[idx*chunkSize:]
				} else {
					chunk = cardDefs[idx*chunkSize : (idx+1)*chunkSize]
				}
				if card.Name != "" && stale {
					database.Omit("name").Save(&chunk)
				} else {
					database.Save(&chunk)
				}
			}
		} else {
			database.CreateInBatches(&cardDefs, chunkSize)
		}
	}

	playerGames := []models.Game{}
	var user models.User
	isCurrUser := cfg.IsUser(userID)
	if userID == cfg.UserId {
		resp, err := games.GetPage(cfg, token, 1)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Latest games request error, see server log for details.")
		}

		if resp.Status() != "200 OK" {
			// usually happens on token expiration
			return c.Redirect().To("/login")
		}

		latestGames := resp.Result().(*games.GameList)
		for _, game := range latestGames.Items {
			model, err := game.ToModel()
			if err != nil {
				fmt.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Card data transformation error, see server log for details.")
			}
			playerGames = append(playerGames, model)
		}

		database.
			Session(&gorm.Session{FullSaveAssociations: true}).
			Clauses(clause.OnConflict{DoNothing: true}).
			Save(&playerGames)

		database.Preload("FactionStats").Preload("Progressions").First(&user, "id = ?", userID)
	} else {
		db.ResetPlayerStats(database, userID)
		resp, err := stats.Get(cfg, userID, token, seasonID)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Ranked stats request error, see server log for details.")
		}

		if resp.Status() != "200 OK" {
			// usually happens on token expiration
			return c.Redirect().To("/login")
		}

		rankedStats := resp.Result().(*stats.RankedStats)
		user, err = rankedStats.ToModel()
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user data.")
		}
		database.Save(&user)
	}

	if user.ID == "" {
		return Render(c, pages.NotFound())
	}
	season := db.GetSeasonName(database, seasonID)
	return Render(c, pages.PlayerProfile(&user, playerGames, season, isCurrUser))
}
