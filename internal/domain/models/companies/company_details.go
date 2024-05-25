package companies

import (
	"github.com/awahids/belajar-go/common"
	"gorm.io/gorm"
)

type CompanyDetails struct {
	common.Base
	FullAddress string    `json:"full_address" gorm:"type:varchar(100);not null"`
	Telpon      int       `json:"telpon" gorm:"type:int;not null"`
	Email       string    `json:"email" gorm:"type:varchar(100);not null"`
	Logo        string    `json:"logo" gorm:"type:varchar(100);not null"`
	BgImage     string    `json:"bg_image" gorm:"type:varchar(100);not null"`
	CompanyId   uint      `json:"company_id"`
	Company     Companies `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}
