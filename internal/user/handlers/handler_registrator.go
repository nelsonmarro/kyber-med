package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/user/repositories"
	"github.com/nelsonmarro/kyber-med/internal/user/services"
)

func RegisterAuthHandlers(router fiber.Router, config *config.Config, db database.Database, middlewares ...fiber.Handler) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHttpHandler(userService, config)

	userRoutes := router.Group("auth", middlewares...)

	userRoutes.Post("/register", userHandler.Register)
	userRoutes.Post("/login", userHandler.Login)
}

func RegisterUserHandlers(router fiber.Router, config *config.Config, db database.Database, middlewares ...fiber.Handler) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHttpHandler(userService, config)

	userRoutes := router.Group("users", middlewares...)

	userRoutes.Get("/:id", userHandler.GetUserByID)
	userRoutes.Patch("/:id", userHandler.UpdateUser)
	userRoutes.Delete("/:id", userHandler.DeleteUser)
}
