package handlers

import "github.com/gofiber/fiber/v3"

type PacientHandler interface {
	GetAllPacients(c *fiber.Ctx) error
}
