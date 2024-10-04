package stats

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"gwentgg/config"
)

func Get(cfg *config.Config, userId string, token string, seasonId string) (*resty.Response, error) {
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

	var stats RankedStats
	url := fmt.Sprintf("https://gwent-rankings.gog.com/ranked_2_0/seasons/%s/users/%s", seasonId, userId)
	return client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(&stats).
		Get(url)
}
