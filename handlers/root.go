package handlers

import (
	"context"

	"awesome.fintech.org/core/clients"
	"awesome.fintech.org/core/constants"
	"go.uber.org/fx"

	"github.com/gofiber/fiber/v3"
	"github.com/rabbitmq/amqp091-go"
)

func NewRootHandler(
	lc fx.Lifecycle,

	env *constants.Env,
	server *fiber.App,
	rabbitmq *clients.RabbitMQ,
	rabbitmqConnection *amqp091.Connection,

	ledgerHandler *LedgerHandler,
	balanceHandler *BalanceHandler,
	transactionHandler *TransactionHandler,
) {
	server.Get("/ledger/", ledgerHandler.List)
	server.Post("/ledger/", ledgerHandler.Create)

	server.Get("/balance/", balanceHandler.List)
	server.Post("/balance/", balanceHandler.Create)
	server.Post("/balance/:balanceId/", balanceHandler.Get)

	server.Get("/transaction/", transactionHandler.List)
	server.Post("/transaction/", transactionHandler.Record)
	server.Post("/transaction/:transactionId/refund/", transactionHandler.Refund)

	registerHandler := func(queue string, callback func(content string), workers int) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go rabbitmq.SubscribeWithWorkers(workers, clients.SubscribeArgs{
					Queue:    queue,
					Callback: callback,
				})

				return nil
			},
			OnStop: func(ctx context.Context) error {
				rabbitmqConnection.Close()

				return nil
			},
		})
	}

	registerHandler(constants.TRANSACTIONS_QUEUE, transactionHandler.HandleRecord, env.TRANSACTION_WORKERS)
}
