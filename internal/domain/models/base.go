package models

type Base struct {
	Id   uint   `gorm:"type:int;not null auto_increment primary_key" `
	UUID string `gorm:"type:string;default:gen_random_uuid()"`
}
