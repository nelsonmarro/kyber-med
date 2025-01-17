package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/pacient/repositories"
	"github.com/nelsonmarro/kyber-med/internal/pacient/services"
)

func RegisterPacientHandlers(router fiber.Router, db database.Database, middlewares ...fiber.Handler) {
	pacientRepo := repositories.NewPacientRepository(db)
	pacientService := services.NewPacientService(pacientRepo)
	pacientHandler := NewPacientHttpHandler(pacientService)

	router.Get("/", pacientHandler.GetPacientsByCursor)
}
