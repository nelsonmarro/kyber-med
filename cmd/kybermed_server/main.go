package main

import (
	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/server"
)

func main() {
	conf := config.LoadConfig("config")

	db := database.NewDatabase(conf)

	server.NewFiberServer(conf, db).Start()
}
