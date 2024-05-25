package companies

import (
	"github.com/awahids/belajar-go/common"
	"gorm.io/gorm"
)

type CompanyEmployees struct {
	common.Base
	Fullname  string    `json:"fullname" gorm:"type:varchar(100);not null"`
	Position  string    `json:"position" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);not null"`
	CompanyId uint      `json:"company_id"`
	Company   Companies `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}
