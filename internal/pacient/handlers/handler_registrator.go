package pacient

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nelsonmarro/kyber-med/internal/database"
	pRepo "github.com/nelsonmarro/kyber-med/internal/pacient/repositories"
	pService "github.com/nelsonmarro/kyber-med/internal/pacient/services"
)

func RegisterPacientHandlers(router fiber.Router, db database.Database, middlewares ...fiber.Handler) {
	pacientRepo := pRepo.NewPacientRepository(db)
	pacientService := pService.NewPacientService(pacientRepo)
	pacientHandler := pRepo.NewPacientHttpHandler(pacientService)

	router.Get("/", pacientHandler.GetPacientsByCursor)
}
