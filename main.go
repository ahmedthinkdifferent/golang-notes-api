package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"noteapp/core/data/db"
	"noteapp/core/env"
	"noteapp/core/http/routes"
	"noteapp/core/util"
)

func main() {
	env.LoadEnv()
	db.Connect()
	app := fiber.New()
	routes.RegisterAppRoutes(app)

	err := app.Listen(util.FormatStr(":%d", env.AppEnvironment.App.Port))
	if err != nil {
		log.Fatal("failed to start app")
	}
}
