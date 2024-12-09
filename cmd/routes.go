package main

import (
	"github.com/thisisvillegas/goTrivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App){
	app.Get("/", handlers.Home)
}