package pacient

import "github.com/gofiber/fiber/v2"

type PacientHandler interface {
	GetPacientsByCursor(c *fiber.Ctx) error
}
