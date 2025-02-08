package main

import (
	"math/rand"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	uConstant "github.com/nelsonmarro/kyber-med/internal/user/constanst"
	uEntities "github.com/nelsonmarro/kyber-med/internal/user/entities"
)

func main() {
	args := os.Args[1:] // Obtener los argumentos pasados por consola

	conf := config.LoadConfig("config")
	db := database.NewDatabase(conf)

	if len(args) > 0 && args[0] == "--migrate" {
		migrateTables(db)
	} else if len(args) > 0 && args[0] == "--seed" {
		usersSeed(db)
	}
}

func migrateTables(db database.Database) {
	db.GetDb().AutoMigrate(&uEntities.User{})
}

func usersSeed(db database.Database) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error encryptar password")
		return
	}

	user := uEntities.User{
		Name:     "Nelson Marro",
		Email:    "nelsonmarro99@gmail.com",
		Role:     uConstant.RoleUser,
		IDCard:   strconv.Itoa(rand.Intn(99999)),
		Password: string(hashedPassword),
	}

	admin := uEntities.User{
		Name:     "Admin",
		Email:    "admin99@gmail.com",
		Role:     uConstant.RoleAdmin,
		IDCard:   strconv.Itoa(rand.Intn(99999)),
		Password: string(hashedPassword),
	}

	db.GetDb().Migrator().DropTable(uEntities.User{})
	db.GetDb().Migrator().CreateTable(uEntities.User{})
	db.GetDb().Create(&user)
	db.GetDb().Create(&admin)
}
