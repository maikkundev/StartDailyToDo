package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name   string `json:"name"`
	IsDone bool   `json:"isDone" gorm:"default:true"`
}
