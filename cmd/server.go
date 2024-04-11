package main

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"awesome.fintech.org/core"
	"awesome.fintech.org/handlers"
)

func startServer(lc fx.Lifecycle) *fiber.App {
	server := core.NewServer()

	handlers.NewRootHandler(server)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Listen(":3000")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})

	return server
}

func ServerRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Database migration utility",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "up",
		Short: "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			fx.New(
				fx.Provide(startServer),
				fx.Invoke(func(*fiber.App) {}),
			).Run()
		},
	})

	return cmd
}
