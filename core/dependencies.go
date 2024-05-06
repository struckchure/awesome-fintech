package core

import (
	"awesome.fintech.org/core/clients"
	"awesome.fintech.org/core/constants"
	"awesome.fintech.org/dao"
	"awesome.fintech.org/handlers"
	"awesome.fintech.org/services"

	"go.uber.org/fx"
)

func SetupDependencies() {
	fx.New(
		fx.Provide(constants.NewEnv),

		fx.Provide(NewDatabaseConnection),
		fx.Provide(clients.NewRabbitMQ),
		fx.Provide(clients.NewRabbitMQConnection),

		fx.Provide(dao.NewLedgerDao),
		fx.Provide(dao.NewBalanceDao),
		fx.Provide(dao.NewTransactionDao),

		fx.Provide(services.NewLedgerService),
		fx.Provide(services.NewBalanceService),
		fx.Provide(services.NewTransactionService),

		fx.Provide(handlers.NewLedgerHandler),
		fx.Provide(handlers.NewBalanceHandler),
		fx.Provide(handlers.NewTransactionHandler),

		fx.Invoke(handlers.NewRootHandler),

		fx.Provide(NewServer),
	).Run()
}
