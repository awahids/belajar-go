package models

import (
	"github.com/awahids/belajar-go/common"
	"github.com/awahids/belajar-go/internal/domain/models/companies"
	"gorm.io/gorm"
)

type Positions struct {
	common.Base
	Title     string              `json:"title" gorm:"unique"`
	Value     string              `json:"value" gorm:"unique"`
	CreatedBy uint                `json:"created_by"`
	Company   companies.Companies `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}
