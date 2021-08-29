package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurliman/anonymous-auth-service/env"
)

func main() {
	env.Load()

	app := fiber.New(fiber.Config{
		AppName: "Anonymous Auth Service",
	})

	app.Listen(":" + env.Get("PORT", "3000"))
}
