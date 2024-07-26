package database

import (
	"fmt"
	config "rojgaarkaro-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {

	dsn := config.Configuration()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection failed. err: -", err)
		return nil, err
	}
	return db, err
}
