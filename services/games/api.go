package games

import (
	"fmt"
	"gwentgg/config"

	"github.com/go-resty/resty/v2"
)

func GetPage(cfg *config.Config, token string, page int) (*resty.Response, error) {
	client := resty.New()
	headers := map[string]string{
		"Authorization":      token,
		"Accept":             "*/*",
		"Accept-Encoding":    "gzip,deflate",
		"User-Agent":         "UnityPlayer/2021.3.15f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)",
		"X-Unity-Version":    "2021.3.15f1",
		"X-Game-BI-Platform": "GOG",
		"X-Game-Version":     "11.10.0",
		"X-Operation":        "4646285001187860130",
		"X-Session":          "7413118591381967144",
	}
	params := map[string]string{
		"page":                fmt.Sprintf("%d", page),
		"_version":            cfg.Version,
		"_data_version":       cfg.DataVersion,
		"_data_version_token": cfg.DataVersionToken,
		"build_region":        cfg.BuildRegion,
	}

	var games GameList
	url := fmt.Sprintf("https://gwent-games-log.gog.com/users/%s/games", cfg.UserId)
	return client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(&games).
		Get(url)
}
