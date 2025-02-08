package pacient

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/common/commondtos"
)

type pacientHttpHandler struct {
	pacientService PacientService
}

func NewPacientHttpHandler(pacientService PacientService) PacientHandler {
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
