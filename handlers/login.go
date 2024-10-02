package handlers

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v3"
	"gwentgg/components"
	"gwentgg/config"
	"gwentgg/db"
	"gwentgg/services"
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

	client := resty.New()
	headers := map[string]string{
		"Authorization":      token,
		"Accept":             "*/*",
		"Accept-Encoding":    "gzip,deflate",
		"User-Agent":         "UnityPlayer/2021.3.15f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)",
		"X-Unity-Version":    "2021.3.15f1",
		"X-Game-BI-Platform": "GOG",
		"X-Game-Version":     "11.10.0",
		"X-Operation":        "4646285001187860127",
		"X-Session":          "7413118591381967144",
	}
	params := map[string]string{
		"_version":      cfg.Version,
		"_data_version": cfg.DataVersion,
		"build_region":  cfg.BuildRegion,
	}

	seasonId := cfg.CurrentSeasonId
	var stats services.RankedStats
	url := fmt.Sprintf("https://gwent-rankings.gog.com/ranked_2_0/seasons/%s/users/%s", seasonId, userId)
	resp, err := client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(&stats).
		Get(url)

	if err != nil {
		fmt.Printf("%s", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ranked stats request error, see server log for details.")
	}

	if resp.Status() != "200 OK" {
		return c.Status(fiber.StatusBadRequest).SendString("Ranked stats request failed, check your credentials.")
	}

	model, err := stats.ToModel()
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
