package models

import (
	"time"

	"gorm.io/gorm"
)

type taskStatus string

const (
	created    taskStatus = "created"
	inprogress taskStatus = "inprogress"
	done       taskStatus = "done"
	failed     taskStatus = "failed"
)

type Todo struct {
	gorm.Model

	ID           uint       `gorm:"autoIncrement; primaryKey"`
	Task         string     `gorm:"type:text;not null"`
	Deescription string     `gorm:"type:text;not null"`
	Status       taskStatus `gorm:"type:task_status; default:created"`
	Due_date     time.Time
	User_id      uint
	Category_id  uint
	Created_by   User         `gorm:"foreignKey:User_id;references:ID"`
	Category     TaskCategory `gorm:"foreignKey:Category_id; references:ID"`
}
