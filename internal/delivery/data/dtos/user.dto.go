package dtos

type CreateUserReq struct {
	Username string   `json:"username" validate:"required"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=8"`
	Role     RoleUuid `json:"role" validate:"required"`
}

type UpdateUserReq struct {
	UUID     string   `json:"uuid" validate:"required"`
	Username string   `json:"username" validate:"required"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=8"`
	RoleUuid RoleUuid `json:"role" validate:"required"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RoleUuid struct {
	RoleUuid string `json:"role_uuid" validate:"required"`
}
