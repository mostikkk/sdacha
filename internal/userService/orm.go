package userService

import (
	task "obuch/internal/taskService"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email string      `json:"email"`
	Pass  string      `json:"pass"`
	Tasks []task.Task `json:"tasks" gorm:"foreignkey:UserID"`
}
