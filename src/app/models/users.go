package models

import (
	"time"
)

type User struct {
	id         uint   `gorm:"autoIncrement; primaryKey"`
	username   string `gorm:"unique"`
	password   string
	created_at time.Time
	updated_at time.Time
}
