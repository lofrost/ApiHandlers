package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task    string `json:"task"`
	IsDone  *bool  `json:"is_done"`
	User_ID uint   `json:"user_id"`
}
