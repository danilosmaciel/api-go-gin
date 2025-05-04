package database

import (
	"github.com/danilosmaciel/api-go-gin/models"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := ""

	if dsn == "" {
		panic("configura o banco de dados em database.go -> Connect")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.City{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.State{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
