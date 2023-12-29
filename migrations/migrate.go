package migrations

import (
	"simple_api/src/app/models"

	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) {

	db.AutoMigrate(
		&models.User{},
		&models.TaskCategory{},
		&models.Todo{},
	)

}
