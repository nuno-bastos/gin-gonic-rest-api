package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "golang-gin-api/config"
	model "golang-gin-api/model"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName)
	db, dbErr := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	db.AutoMigrate(&model.Tags{})

	return db, dbErr
}
