package models

type Role struct {
	ID   uint     `json:"id"`
	Name RoleType `json:"name"`
}
