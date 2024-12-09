package database

import (

	"fmt"
	"log"
	"os"

	"github.com/thisisvillegas/goTrivia/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb(){
	log.Println("os.GetEnv(DB_USER)", os.GetEnv("DB_USER"))
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Central", 
		os.GetEnv("DB_USER"),
		os.GetEnv("DB_PASSWORD"),
		os.GetEnv("DB_NAME"),
	)


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)
	})

	if err != nil {
		log.Fatal("Failed to connect ot database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.logMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db:db,
	}
}