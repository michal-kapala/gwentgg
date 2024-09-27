package main

import (
	"fmt"

	"gwentgg/components"
	"gwentgg/config"
	"gwentgg/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		fmt.Println("Failed to load configuration:", err)
		return
	}

	app := fiber.New()
	app.Use(config.Add(cfg))
	app.Get("/login", handlers.LoginHandler)
	app.Post("/login", handlers.LoginSubmitHandler)
	log.Fatal(app.Listen(":3000"))
}

func NotFoundMiddleware(c fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return handlers.Render(c, components.NotFound())
}
