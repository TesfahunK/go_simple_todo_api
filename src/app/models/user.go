package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"autoIncrement; primaryKey"`
	Username string `gorm:"unique"`
	Password string
}
