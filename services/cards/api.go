package cards

import (
	"gwentgg/config"

	"github.com/go-resty/resty/v2"
)

func Get(cfg *config.Config, token string) (*resty.Response, error) {
	client := resty.New()
	headers := map[string]string{
		"Authorization":      token,
		"Accept":             "*/*",
		"Accept-Encoding":    "gzip,deflate",
		"User-Agent":         "UnityPlayer/2021.3.15f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)",
		"X-Unity-Version":    "2021.3.15f1",
		"X-Game-BI-Platform": "GOG",
		"X-Game-Version":     "11.10.0",
		"X-Operation":        "4646285001187860129",
		"X-Session":          "7413118591381967144",
	}
	params := map[string]string{
		"_version":            cfg.Version,
		"_data_version":       cfg.DataVersion,
		"_data_version_token": cfg.DataVersionToken,
		"build_region":        cfg.BuildRegion,
	}

	var cards CardDefList
	url := "https://gwent-deck.gog.com/card_definitions"
	return client.R().
		SetHeaders(headers).
		SetQueryParams(params).
		SetResult(&cards).
		Get(url)
}
