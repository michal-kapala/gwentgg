package handlers

import (
	"fmt"
	"strings"
	"time"

	"gwentgg/components"
	"gwentgg/config"
	"gwentgg/db"
	"gwentgg/db/models"
	"gwentgg/services/seasons"
	"gwentgg/services/stats"

	"github.com/gofiber/fiber/v3"
)

func LoginHandler(c fiber.Ctx) error {
	userId := config.Get(c).UserId
	if userId == "" {
		return Render(c, components.UserIdMissing())
	}
	return Render(c, components.LoginForm(userId))
}

func LoginSubmitHandler(c fiber.Ctx) error {
	userId := c.FormValue("user_id")
	token := c.FormValue("access_token")
	cfg := config.Get(c)
	database := db.Get(c)
	db.ResetPlayerStats(database, userId)

	if !strings.HasPrefix(token, "Bearer ") {
		token = fmt.Sprintf("Bearer %s", token)
	}

	var season models.Season
	database.Where("? BETWEEN date_starts AND date_ends", time.Now()).First(&season)
	if season.ID == "" {
		page := 1
		resp, err := seasons.Get(cfg, userId, token, page)
		if err != nil {
			fmt.Printf("%s", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Ranked seasons request error, see server log for details.")
		}
		seasonList := resp.Result().(*seasons.SeasonList)
		seasonModels, err := seasons.ToModels(seasonList.Items)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Season data transformation error, see server log for details.")
		}
		database.Create(&seasonModels)

		database.Where("? BETWEEN date_starts AND date_ends", time.Now()).First(&season)
		if season.ID == "" {
			msg := "No current season found."
			fmt.Printf("%s\n", msg)
			return c.Status(fiber.StatusInternalServerError).SendString(msg)
		}
	}
	fmt.Printf("Current season: %s\n", season.Name)
	resp, err := stats.Get(cfg, userId, token, season.ID)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ranked stats request error, see server log for details.")
	}

	if resp.Status() != "200 OK" {
		return c.Status(fiber.StatusBadRequest).SendString("Ranked stats request failed, check your credentials.")
	}

	rankedStats := resp.Result().(*stats.RankedStats)
	statsModel, err := rankedStats.ToModel()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user data.")
	}
	database.Save(&statsModel)

	c.Cookie(&fiber.Cookie{
		Name:  "access_token",
		Value: token,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "token_expired",
		Value: "false",
	})
	return c.Redirect().To(fmt.Sprintf("/player/%s", userId))
}
