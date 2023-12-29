package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpendbConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=todo port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	return db, err
}
