package models

import "time"

type TaskCategory struct {
	id         uint   `gorm:"autoIncrement;primaryKey"`
	name       string `gorm:"unique"`
	created_at time.Time
}
