package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint      `gorm:"type:int;primary_key"`
	UUID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Title  string    `json:"title" gorm:"type:varchar(100);not null"`
	Author string    `json:"author" gorm:"type:varchar(100);not null"`
	Year   int       `json:"year" gorm:"type:int;not null"`
	gorm.Model
}
