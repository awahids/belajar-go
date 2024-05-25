package companies

import (
	"github.com/awahids/belajar-go/common"
	"gorm.io/gorm"
)

type CompanyLocations struct {
	common.Base
	Latitude  string    `json:"latitude" gorm:"type:varchar(100);not null"`
	Longitude string    `json:"longitude" gorm:"type:varchar(100);not null"`
	Distance  int       `json:"distance" gorm:"type:int;not null"`
	CompanyId uint      `json:"company_id"`
	Company   Companies `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}
