package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"

	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
)

type fiberServer struct {
	app  *fiber.App
	db   database.Database
	conf *config.Config
}

func NewFiberServer(conf *config.Config, db database.Database) Server {
	fiberApp := fiber.New()
	log.SetLevel(log.LevelDebug)

	return &fiberServer{
		app:  fiberApp,
		db:   db,
		conf: conf,
	}
}

func (s *fiberServer) Start() {
	s.app.Use(recover.New())
	s.app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/Guayaquil",
	}))

	// Provide a minimal config for liveness check
	s.app.Get(healthcheck.DefaultLivenessEndpoint, healthcheck.NewHealthChecker())
	// Provide a minimal config for readiness check
	s.app.Get(healthcheck.DefaultReadinessEndpoint, healthcheck.NewHealthChecker())
	// Provide a minimal config for startup check
	s.app.Get(healthcheck.DefaultStartupEndpoint, healthcheck.NewHealthChecker())

	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(s.app.Listen(":3000"))
}
