package database

import (
	"fmt"
	"log"
	"task-5-vix-btpns-SofyanEgiLesmana/helpers"
	"task-5-vix-btpns-SofyanEgiLesmana/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	helpers.LoadEnv(".env")
	host := helpers.GetEnv("DB_HOST", "localhost")
	user := helpers.GetEnv("DB_USER", "postgres")
	password := helpers.GetEnv("DB_PASSWORD", "postgres")
	dbName := helpers.GetEnv("DB_NAME", "default")
	dbPort := helpers.GetEnv("DB_PORT", "5432")

	dbURI := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbName,
		dbPort)

	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database:", err.Error())
	}
	fmt.Println("Success connected to database")
}

func MigateDB() {
	db.Debug().AutoMigrate(
		models.User{},
		models.Photo{},
	)
}

func GetDB() *gorm.DB {
	return db
}
