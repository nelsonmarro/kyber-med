package pacient

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/common/commondtos"
	pDtos "github.com/nelsonmarro/kyber-med/internal/pacient/dtos"
	pService "github.com/nelsonmarro/kyber-med/internal/pacient/services"
)

type pacientHttpHandler struct {
	pacientService pService.PacientService
}

func NewPacientHttpHandler(pacientService pService.PacientService) PacientHandler {
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
	resp := commondtos.ResponseDTO[pDtos.PacientDto]{
		Success:    true,
		Data:       pacientes,
		Pagination: pagination,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
