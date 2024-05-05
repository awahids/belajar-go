package models

import (
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     uint   `json:"role_id"`
	gorm.Model
}

type RoleType string

const (
	AdminRole RoleType = "admin"
	UserRole  RoleType = "user"
)
