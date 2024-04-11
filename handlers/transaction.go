package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"

	"awesome.fintech.org/dto"
	"awesome.fintech.org/services"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func (h *TransactionHandler) Record(c fiber.Ctx) error {
	var dto dto.CreateTransactionDto
	err := json.Unmarshal(c.Body(), &dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	transaction, err := h.transactionService.Record(dto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(transaction)
}

func (h *TransactionHandler) Refund(c fiber.Ctx) error {
	transactionId := c.Params("transactionId")
	transaction, err := h.transactionService.Refund(dto.GetTransactionDto{Id: transactionId})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(transaction)
}

func NewTransactionHandler() *TransactionHandler {
	transactionService := services.NewTransactionService()

	return &TransactionHandler{transactionService: transactionService}
}
