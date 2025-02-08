package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"

	commonentities "github.com/nelsonmarro/kyber-med/common/commonentities"
	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	pEntities "github.com/nelsonmarro/kyber-med/internal/pacient/entities"
	uEntities "github.com/nelsonmarro/kyber-med/internal/user/entities"
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
	db.GetDb().AutoMigrate(&pEntities.Pacient{})
}

func pacientsSeed(db database.Database) {
	date, err := time.Parse("2006-01-02", "1999-01-09")
	if err != nil {
		log.Error("error al parsear la fecha")
		return
	}

	var user uEntities.User
	_ = db.GetDb().First(&user)

	date = date.Add(-time.Duration(21-1) * time.Hour)
	pacient := pEntities.Pacient{
		BaseEntity:            commonentities.BaseEntity{},
		FirstName:             fmt.Sprintf("Paciente %d", 1),
		LastName:              fmt.Sprintf("Last %d", 1),
		Email:                 fmt.Sprintf("nelsonmarro%d@gmail.com", 1),
		IDCard:                strconv.Itoa(rand.Intn(99999)),
		PhoneNumber:           "0985134196",
		DateOfBirth:           date,
		Gender:                "Masculino",
		Address:               "Quito",
		EmergencyContactName:  "Alieen Torres",
		EmergencyContactPhone: "0999079590",
		User:                  user,
	}

	db.GetDb().Migrator().DropTable(&pEntities.Pacient{})
	db.GetDb().Create(pacient)
}
