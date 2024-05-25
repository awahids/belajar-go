package models

import (
	"github.com/awahids/belajar-go/common"
	"gorm.io/gorm"
)

type User struct {
	common.Base
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
	Role     Roles  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}

type RoleType string

const (
	AdminRole RoleType = "admin"
	UserRole  RoleType = "user"
)
