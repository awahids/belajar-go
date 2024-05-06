package models

import "gorm.io/gorm"

type Roles struct {
	Base
	Title RoleType `json:"title" gorm:"unique"`
	Value string   `json:"value" gorm:"unique"`
	gorm.Model
}
