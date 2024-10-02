package main

import (
	"fmt"

	"gwentgg/components"
	"gwentgg/config"
	"gwentgg/db"
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

	database, err := db.Init("gwent.db")
	if database == nil {
		fmt.Println(err)
		return
	}

	err = db.AutoMigrate(database)
	if err != nil {
		fmt.Println("DB migration failed:", err)
		return
	}

	app := fiber.New()
	app.Use(config.Add(cfg))
	app.Use(db.Add(database))
	app.Get("/", handlers.IndexHandler)
	app.Get("/login", handlers.LoginHandler)
	app.Post("/login", handlers.LoginSubmitHandler)
	log.Fatal(app.Listen(":3000"))
}

func NotFoundMiddleware(c fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return handlers.Render(c, components.NotFound())
}
