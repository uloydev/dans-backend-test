package main

import (
	"dans-backend-test/config"
	"dans-backend-test/db"
	"dans-backend-test/initialize"

	_ "dans-backend-test/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title                       Dans Backnd test
// @version                     1.0
// @description                 Dans Backnd test
// @contact.name                wahyu miftahul aflah
// @contact.email               wahyumiftahul7@gmail.com
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath                    /
func main() {
	conf := config.New()
	dbConn := db.NewGormConnection(conf)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	v1 := app.Group("/v1")

	initialize.RunInitFunctions(v1.(*fiber.Group), dbConn)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "http://localhost:6991/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	app.Listen(":6991")
}
