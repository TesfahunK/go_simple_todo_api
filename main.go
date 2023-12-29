package main

import (
	"simple_api/migrations"
	"simple_api/src/app/models"
	"simple_api/src/config"
	"simple_api/src/config/logger"

	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func main() {

	intialSetups()
	user := models.User{Username: "adam", Password: "password"}

	Database.Create(&user)

}

func setupDb() error {

	var err error

	Database, err = config.OpendbConnection()

	if err != nil {
		return err
	}
	logger.Debug("Database Connected")

	migrations.MigrateTables(Database)

	return nil
}

func intialSetups() {
	// setup global logger here
	logger.SetupLogger()
	setupDb()

}
