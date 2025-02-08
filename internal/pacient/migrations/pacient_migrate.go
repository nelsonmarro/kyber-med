package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"

	commonentities "github.com/nelsonmarro/kyber-med/common/commonentities"
	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
)

func main() {
	args := os.Args[1:] // Obtener los argumentos pasados por consola

	conf := config.LoadConfig("config")
	db := database.NewDatabase(conf)

	if len(args) > 0 && args[0] == "--migrate" {
		migrateTables(db)
	} else if len(args) > 0 && args[0] == "--seed" {
		pacientsSeed(db)
	}
}

func migrateTables(db database.Database) {
	db.GetDb().AutoMigrate(&Pacient{})
}

func pacientsSeed(db database.Database) {
	date, err := time.Parse("2006-01-02", "1999-01-09")
	if err != nil {
		log.Error("error al parsear la fecha")
		return
	}

	pacients := make([]Pacient, 0)

	for i := 0; i < 10; i++ {
		date = date.Add(-time.Duration(21-i) * time.Hour)
		pacients = append(pacients, Pacient{
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
			User:                  Pser,
		})
	}

	db.GetDb().AutoMigrate(&Pacient{})

	var pacient Pacient
	result := db.GetDb().First(&pacient)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.GetDb().CreateInBatches(pacients, 10)
	}
}
