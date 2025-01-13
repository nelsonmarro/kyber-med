package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/internal/user/dtos"
	"github.com/nelsonmarro/kyber-med/internal/user/services"
)

type authHttpHandler struct {
	userService services.UserService
}

func NewAuthHttpHandler(userService services.UserService) AuthHandler {
	return &authHttpHandler{userService: userService}
}

func (h *authHttpHandler) Register(c *fiber.Ctx) error {
	var req dtos.UserRegisterDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := h.userService.RegisterUser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
	})
}
