package usersorm

import (
	"log"
	"starter-pack-api/internal/config"
	"starter-pack-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(configDb config.DatabaseConn) {
	connectionString := "postgresql://" + configDb.Username +
		":" + configDb.Password +
		"@" + configDb.Host + ":" + configDb.Port +
		"/" + configDb.DatabaseName + "?sslmode=disable"
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed...")
}
