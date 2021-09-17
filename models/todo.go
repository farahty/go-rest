package models

import "github.com/google/uuid"

type Todo struct {
	Base
	Name      string `json:"name"`
	Completed bool   `json:"completed" gorm:"default:false"`

	User   User      `json:"user"`
	UserID uuid.UUID `json:"userId"`
}
