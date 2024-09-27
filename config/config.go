package config

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3"
)

type Config struct {
	DataVersion				string	`json:"_data_version"`
	DataVersionToken	string	`json:"_data_version_token"`
	Version						string	`json:"_version"`
	BuildRegion				string	`json:"build_region"`
	CurrentSeasonId		string	`json:"current_season_id"`
	UserId						string	`json:"user_id"`
}

func Load(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
			return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
			return nil, err
	}

	return &config, nil
}

func Add(config *Config) fiber.Handler {
	return func(c fiber.Ctx) error {
		c.Locals("config", config)
		return c.Next()
	}
}

func Get(c fiber.Ctx) *Config {
	return c.Locals("config").(*Config)
}
