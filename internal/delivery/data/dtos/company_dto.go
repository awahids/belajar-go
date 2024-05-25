package dtos

type CreateCompanyReq struct {
	Name string `validate:"required" json:"name"`
}
