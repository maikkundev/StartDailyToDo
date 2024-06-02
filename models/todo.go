package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"isDone" gorm:"default:true"`
}
