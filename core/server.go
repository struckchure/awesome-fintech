package core

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle) *fiber.App {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(":3000")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
