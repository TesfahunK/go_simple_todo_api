package models

import "time"

type TaskCategory struct {
	ID         uint   `gorm:"autoIncrement;primaryKey"`
	name       string `gorm:"unique"`
	created_at time.Time
}
