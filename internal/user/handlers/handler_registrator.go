package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	uRepo "github.com/nelsonmarro/kyber-med/internal/user/repositories"
	uServices "github.com/nelsonmarro/kyber-med/internal/user/services"
)

func RegisterAuthHandlers(router fiber.Router, config *config.Config, db database.Database, middlewares ...fiber.Handler) {
	userRepo := uRepo.NewUserRepository(db)
	userService := uServices.NewUserService(userRepo)
	userHandler := NewUserHttpHandler(userService, config)

	router.Post("/register", userHandler.Register)
	router.Post("/login", userHandler.Login)
}

func RegisterUserHandlers(router fiber.Router, config *config.Config, db database.Database, middlewares ...fiber.Handler) {
	userRepo := uRepo.NewUserRepository(db)
	userService := uServices.NewUserService(userRepo)
	userHandler := NewUserHttpHandler(userService, config)

	router.Get("/:id", userHandler.GetUserByID)
	router.Patch("/:id", userHandler.UpdateUser)
	router.Delete("/:id", userHandler.DeleteUser)
}
