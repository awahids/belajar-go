package companies

import (
	"github.com/awahids/belajar-go/common"
	"gorm.io/gorm"
)

type Companies struct {
	common.Base
	Name string `json:"name" gorm:"unique not null"`
	Code string `json:"code" gorm:"unique not null"`
	gorm.Model
}
