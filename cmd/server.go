package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"awesome.fintech.org/core"
	"awesome.fintech.org/dao"
	"awesome.fintech.org/handlers"
	"awesome.fintech.org/services"
)

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
				fx.Provide(core.NewDatabaseConnection),

				fx.Provide(dao.NewLedgerDao),
				fx.Provide(services.NewLedgerService),
				fx.Provide(handlers.NewLedgerHandler),

				fx.Provide(dao.NewBalanceDao),
				fx.Provide(services.NewBalanceService),
				fx.Provide(handlers.NewBalanceHandler),

				fx.Provide(dao.NewTransactionDao),
				fx.Provide(services.NewTransactionService),
				fx.Provide(handlers.NewTransactionHandler),

				fx.Provide(core.NewServer),

				fx.Invoke(handlers.NewRootHandler),
				fx.Invoke(func(*fiber.App) {}),
			).Run()
		},
	})

	return cmd
}
