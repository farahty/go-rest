package models

type Todo struct {
	Base
	Name      string `json:"name"`
	Completed bool   `json:"completed" gorm:"default:false"`
}
