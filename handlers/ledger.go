package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"

	"awesome.fintech.org/dto"
	"awesome.fintech.org/services"
)

type LedgerHandler struct {
	ledgerService *services.LedgerService
}

func (h *LedgerHandler) List(c fiber.Ctx) error {
	ledgers, err := h.ledgerService.List(dto.ListLedgerDto{})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(ledgers)
}

func (h *LedgerHandler) Create(c fiber.Ctx) error {
	var dto dto.CreateLedgerDto
	err := json.Unmarshal(c.Body(), &dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ledgers, err := h.ledgerService.Create(dto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(ledgers)
}

func NewLedgerHandler() *LedgerHandler {
	ledgerService := services.NewLedgerService()

	return &LedgerHandler{ledgerService: ledgerService}
}
