package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/common/commondtos"
	"github.com/nelsonmarro/kyber-med/internal/pacient/dtos"
	"github.com/nelsonmarro/kyber-med/internal/pacient/services"
)

type pacientHttpHandler struct {
	pacientService services.PacientService
}

func NewPacientHttpHandler(pacientService services.PacientService) PacientHandler {
	return &pacientHttpHandler{
		pacientService: pacientService,
	}
}

func (h *pacientHttpHandler) GetPacientsByCursor(c *fiber.Ctx) error {
	perPageStr := c.Query("per_page", "10")
	sortOrder := c.Query("sort_order", "asc")
	cursor := c.Query("cursor", "")

	limit, err := strconv.Atoi(perPageStr)
	if err != nil {
		limit = 10
	}

	pacientes, pagination, err := h.pacientService.GetPacientsByCursor(cursor, limit, sortOrder)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Construir respuesta strongly typed
	resp := commondtos.ResponseDTO[dtos.PacientDto]{
		Success:    true,
		Data:       pacientes,
		Pagination: pagination,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
