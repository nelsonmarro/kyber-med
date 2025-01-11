package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2/middleware/healthcheck"

	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/pacient/handlers"
	"github.com/nelsonmarro/kyber-med/internal/server/middlewares"
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

	s.app.Use(healthcheck.New())

	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	handlers.RegisterPacientHandlers(v1, s.db, middlewares.NewJwtMiddleware(s.conf.Jwt.Key))

	log.Fatal(s.app.Listen(":3000"))
}
