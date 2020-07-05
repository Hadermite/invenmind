package main

import (
	"github.com/Hadermite/invenmind/database"
	"github.com/Hadermite/invenmind/middleware"
	"github.com/Hadermite/invenmind/routes"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	var app = fiber.New()
	database.Initialize()

	app.Use(middleware.ValidateAuth)
	routes.User(app.Group("user"))

	app.Listen(3000)

	defer database.Connection.Close()
}
