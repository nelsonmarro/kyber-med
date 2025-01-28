package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/nelsonmarro/kyber-med/common/commonhelpers"
	"github.com/nelsonmarro/kyber-med/common/commonhelpers/jwthelpers"
	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/user/dtos"
	"github.com/nelsonmarro/kyber-med/internal/user/services"
)

type userHttpHandler struct {
	userService services.UserService
	conf        *config.Config
}

func NewUserHttpHandler(userService services.UserService, conf *config.Config) UserHandler {
	return &userHttpHandler{userService: userService, conf: conf}
}

func (h *userHttpHandler) Register(c *fiber.Ctx) error {
	var req dtos.UserRegisterDTO
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

func (h *userHttpHandler) Login(c *fiber.Ctx) error {
	var loginInput dtos.UserLoginDTO
	if err := c.BodyParser(&loginInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "errors": err.Error()})
	}

	user, pass, err := new(dtos.UserDTO), *new(string), *new(error)

	if commonhelpers.IsEmailValid(loginInput.Identity) {
		user, pass, err = h.userService.GetUserWithPasswordByEmail(loginInput.Identity)
	} else {
		user, pass, err = h.userService.GetUserWithPassswordByIDCard(loginInput.Identity)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal Server Error", "data": err})
	} else if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": err})
	}

	if !commonhelpers.CheckPasswordHash(pass, loginInput.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": nil})
	}

	token, err := jwthelpers.GenerateToken(user.ID, user.Name, user.Email, string(user.Role), h.conf.Jwt.Key)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Success login", "data": dtos.LoginResponseDTO{
		Token: token,
	}})
}

func (h *userHttpHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userService.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "errors": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Success login", "data": user})
}

func (h *userHttpHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)
	var userUpdate dtos.UserUpdateDTO

	if !jwthelpers.ValidToken(token, id) {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	if err := c.BodyParser(&userUpdate); err != nil {
		return c.Status(fiber.StatusInsufficientStorage).JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err})
	}

	err := h.userService.UpdateUser(userUpdate, id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "A problem occurs while updating the user", "errors": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Updated", "data": nil})
}

func (h *userHttpHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !jwthelpers.ValidToken(token, id) {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	err := h.userService.DeleteUser(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "A problem occurs while deleting the user", "errors": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Deleted", "data": nil})
}
