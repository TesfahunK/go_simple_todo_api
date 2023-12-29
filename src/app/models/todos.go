package models

import "time"

type taskStatus string

const (
	created    taskStatus = "created"
	inprogress taskStatus = "inprogress"
	done       taskStatus = "done"
	failed     taskStatus = "failed"
)

type Todo struct {
	id          uint       `gorm:"autoIncrement; primaryKey"`
	task        string     `gorm:"type:text;not null"`
	description string     `gorm:"type:text;not null"`
	status      taskStatus `gorm:"type:task_status; default:created"`
	due_date    time.Time
	created_at  time.Time
	user_id     uint
	category_id uint
	created_by  User         `gorm:"foreignKey:user_id;references:id"`
	category    TaskCategory `gorm:"foreignKey:category_id; references:id"`
}
