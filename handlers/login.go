package handlers

import (
	"fmt"
	"gwentgg/config"
	"gwentgg/components"
	"gwentgg/services"
	"github.com/go-resty/resty/v2"
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
	season_id := config.Get(c).CurrentSeasonId
	
	client := resty.New()
	headers := map[string]string {
		"Authorization": fmt.Sprintf("Bearer %s", token),
		"Accept": "*/*",
		"Accept-Encoding": "gzip,deflate",
		"User-Agent": "UnityPlayer/2021.3.15f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)",
		"X-Unity-Version": "2021.3.15f1",
		"X-Game-BI-Platform": "GOG",
		"X-Game-Version": "11.10.0",
		"X-Operation": "4646285001187860127",
		"X-Session": "7413118591381967144",
	}

	var stats services.RankedStats
	url := fmt.Sprintf("https://gwent-rankings.gog.com/ranked_2_0/seasons/%s/users/%s", season_id, userId)
	resp, err := client.R().
		SetHeaders(headers).
		SetResult(&stats).
		Get(url)

	if err != nil {
		fmt.Printf("%s", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Ranked stats request error, see server log for details.")
	}
	
	if resp.Status() != "200 OK" {
		return c.Status(fiber.StatusBadRequest).SendString("Ranked stats request failed, check your credentials.")
	}

	// save user stats

	c.Cookie(&fiber.Cookie{
		Name: "access_token",
		Value: token,
	})
	return c.Status(fiber.StatusOK).SendString("Authentication successful.")
	//return c.Redirect().Route(fmt.Sprintf("/%s", userId))
}
