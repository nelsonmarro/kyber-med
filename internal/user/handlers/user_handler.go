package user

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
}
