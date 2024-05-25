package dtos

type CreateRoleReq struct {
	Title string `validate:"required" json:"title"`
	Value string `validate:"required" json:"value"`
}
