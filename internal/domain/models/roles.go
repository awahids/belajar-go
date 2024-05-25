package models

import (
	"github.com/awahids/belajar-go/common"
	"gorm.io/gorm"
)

type Roles struct {
	common.Base
	Title RoleType `json:"title" gorm:"unique"`
	Value string   `json:"value" gorm:"unique"`
	gorm.Model
}
