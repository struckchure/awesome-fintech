package handlers

import (
	"github.com/gofiber/fiber/v3"
)

type RootHandler struct{}

func NewRootHandler(server *fiber.App) RootHandler {
	ledgerHandler := NewLedgerHandler()
	balanceHandler := NewBalanceHandler()
	transactionHandler := NewTransactionHandler()

	server.Get("/ledger/", ledgerHandler.List)
	server.Post("/ledger/", ledgerHandler.Create)

	server.Post("/balance/", balanceHandler.Create)
	server.Post("/balance/:balanceId/", balanceHandler.Get)

	server.Post("/transaction/", transactionHandler.Record)
	server.Post("/transaction/:transactionId/refund/", transactionHandler.Refund)

	return RootHandler{}
}
