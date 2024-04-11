package core

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func NewServer() *fiber.App {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New())

	return app
}
