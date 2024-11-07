package database

import (
	"log"

	"fahmi-wallet/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect() {
	var err error
	// Connect to PostgreSQL
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=voca password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database")

	// Migrate schema
	DB.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{}, &models.TransactionType{}, &models.Product{})
}
