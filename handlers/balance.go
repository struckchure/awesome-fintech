package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"

	"awesome.fintech.org/dto"
	"awesome.fintech.org/services"
)

type BalanceHandler struct {
	balanceService *services.BalanceService
}

func (h *BalanceHandler) List(c fiber.Ctx) error {
	balance, err := h.balanceService.List(dto.ListBalanceDto{})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(balance)
}

func (h *BalanceHandler) Create(c fiber.Ctx) error {
	var dto dto.CreateBalanceDto
	err := json.Unmarshal(c.Body(), &dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	balance, err := h.balanceService.Create(dto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(balance)
}

func (h *BalanceHandler) Get(c fiber.Ctx) error {
	balanceId := c.Params("balanceId")
	balance, err := h.balanceService.Get(dto.GetBalanceDto{Id: balanceId})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(balance)
}

func NewBalanceHandler(balanceService *services.BalanceService) *BalanceHandler {
	return &BalanceHandler{balanceService: balanceService}
}
