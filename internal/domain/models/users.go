package models

import (
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
	Role     Roles  `gorm:"foreignKey:RoleId"`
	gorm.Model
}

type RoleType string

const (
	AdminRole RoleType = "admin"
	UserRole  RoleType = "user"
)
