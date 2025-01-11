package main

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/nelsonmarro/kyber-med/config"
	"github.com/nelsonmarro/kyber-med/internal/database"
	"github.com/nelsonmarro/kyber-med/internal/user/constanst"
	"github.com/nelsonmarro/kyber-med/internal/user/entities"
)

func main() {
	conf := config.LoadConfig("config")
	db := database.NewDatabase(conf)

	pacientMigrate(db)
}

func pacientMigrate(db database.Database) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error encryptar password")
		return
	}

	user := entities.User{
		Role:     constanst.RoleUser,
		IDCard:   strconv.Itoa(rand.Intn(99999)),
		Password: string(hashedPassword),
	}

	admin := entities.User{
		Role:     constanst.RoleAdmin,
		IDCard:   strconv.Itoa(rand.Intn(99999)),
		Password: string(hashedPassword),
	}

	db.GetDb().AutoMigrate(&entities.User{})

	var userDb entities.User
	result := db.GetDb().First(&userDb)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.GetDb().Create(&user)
		db.GetDb().Create(&admin)
	}
}
