package config

import (
	"simple_api/migrations"
	"simple_api/src/config/logger"

	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func setupDb() error {

	var err error

	Database, err = OpendbConnection()

	if err != nil {
		return err
	}
	logger.Debug("Database Connected")

	migrations.MigrateTables(Database)

	return nil
}

func InitialSetups() {
	// setup global logger here
	logger.SetupLogger()
	setupDb()

}
