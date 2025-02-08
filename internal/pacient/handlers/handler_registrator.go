package pacient

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/internal/database"
)

func RegisterPacientHandlers(router fiber.Router, db database.Database, middlewares ...fiber.Handler) {
	pacientRepo := NewPacientRepository(db)
	pacientService := NewPacientService(pacientRepo)
	pacientHandler := NewPacientHttpHandler(pacientService)

	router.Get("/", pacientHandler.GetPacientsByCursor)
}
