package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gwentgg/components"
	"gwentgg/config"
	"gwentgg/db"
	"gwentgg/services/stats"
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

	seasonId := cfg.CurrentSeasonId
	resp, err := stats.Get(cfg, userId, token, seasonId)

	if err != nil {
		fmt.Printf("%s", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ranked stats request error, see server log for details.")
	}

	if resp.Status() != "200 OK" {
		return c.Status(fiber.StatusBadRequest).SendString("Ranked stats request failed, check your credentials.")
	}

	rankedStats := resp.Result().(*stats.RankedStats)
	model, err := rankedStats.ToModel()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user data.")
	}
	database.Save(&model)

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
