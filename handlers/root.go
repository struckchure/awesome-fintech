package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func NewRootHandler(
	server *fiber.App,

	ledgerHandler *LedgerHandler,
	balanceHandler *BalanceHandler,
	transactionHandler *TransactionHandler,
) {
	server.Get("/ledger/", ledgerHandler.List)
	server.Post("/ledger/", ledgerHandler.Create)

	server.Get("/balance/", balanceHandler.List)
	server.Post("/balance/", balanceHandler.Create)
	server.Post("/balance/:balanceId/", balanceHandler.Get)

	server.Post("/transaction/", transactionHandler.Record)
	server.Post("/transaction/:transactionId/refund/", transactionHandler.Refund)
}
