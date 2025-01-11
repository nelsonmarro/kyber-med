package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"

	commonentities "github.com/nelsonmarro/kyber-med/common/commonentities"
	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/pacient/entities"
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

	pacients := make([]entities.Pacient, 0)

	for i := 0; i < 10; i++ {
		date = date.Add(-time.Duration(21-i) * time.Hour)
		pacients = append(pacients, entities.Pacient{
			BaseEntity:            commonentities.BaseEntity{},
			FirstName:             fmt.Sprintf("Paciente %d", i),
			LastName:              fmt.Sprintf("Last %d", i),
			Email:                 fmt.Sprintf("nelsonmarro%d@gmail.com", i),
			IDCard:                strconv.Itoa(rand.Intn(99999)),
			PhoneNumber:           "0985134196",
			DateOfBirth:           date,
			Gender:                "Masculino",
			Address:               "Quito",
			EmergencyContactName:  "Alieen Torres",
			EmergencyContactPhone: "0999079590",
		})
	}

	db.GetDb().AutoMigrate(&entities.Pacient{})

	var pacient entities.Pacient
	result := db.GetDb().First(&pacient)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.GetDb().CreateInBatches(pacients, 10)
	}
}
