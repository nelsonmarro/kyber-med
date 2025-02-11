package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2/middleware/healthcheck"

	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	pacienthandlers "github.com/nelsonmarro/kyber-med/internal/pacient/handlers"
	"github.com/nelsonmarro/kyber-med/internal/server/middlewares"
	userhandlers "github.com/nelsonmarro/kyber-med/internal/user/handlers"
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

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8100",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	jwt := middlewares.NewJwtMiddleware(s.conf.Jwt.Key)

	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	userhandlers.RegisterAuthHandlers(auth, s.conf, s.db)

	users := v1.Group("/users", jwt)
	userhandlers.RegisterUserHandlers(users, s.conf, s.db)

	pacients := v1.Group("/pacients", jwt)
	pacienthandlers.RegisterPacientHandlers(pacients, s.db)

	// s.app.Use(func(c *fiber.Ctx) error {
	// 	return c.SendStatus(404) // => 404 "Not Found"
	// })

	log.Fatal(s.app.Listen(":3000"))
}
