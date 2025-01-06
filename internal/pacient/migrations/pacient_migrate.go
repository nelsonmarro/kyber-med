package main

import (
	"time"

	"github.com/gofiber/fiber/v3/log"

	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/pacient/entities"
	shared "github.com/nelsonmarro/kyber-med/internal/shared/entities"
)

func main() {
	conf := config.LoadConfig("config")
	db := database.NewDatabase(conf)

	pacientMigrate(db)
}

func pacientMigrate(db database.Database) {
	date, err := time.Parse("2006-01-02", "1999-01-09")
	if err != nil {
		log.Error("error al parsear la fecha")
		return
	}

	db.GetDb().Migrator().CreateTable(&entities.Pacient{})
	db.GetDb().CreateInBatches([]entities.Pacient{
		{
			BaseEntity:            shared.BaseEntity{},
			FirstName:             "Nelson",
			LastName:              "Marro",
			Email:                 "nelsonmarro99@gmail.com",
			IDCard:                "1757078579",
			PhoneNumber:           "0985134196",
			DateOfBirth:           date,
			Gender:                "Masculino",
			Address:               "Quito",
			EmergencyContactName:  "Alieen Torres",
			EmergencyContactPhone: "0999079590",
		},
		{
			BaseEntity:            shared.BaseEntity{},
			FirstName:             "Maria",
			LastName:              "José",
			Email:                 "majitopc1998@gmail.com",
			IDCard:                "1757078573",
			PhoneNumber:           "0985134192",
			DateOfBirth:           date,
			Gender:                "Femenino",
			Address:               "Latacunga",
			EmergencyContactName:  "Alieen Torres",
			EmergencyContactPhone: "0999079590",
		},
	}, 2)
}
