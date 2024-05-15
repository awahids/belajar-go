package response

type UserResponse struct {
	Id       int          `json:"id"`
	UUID     string       `json:"uuid"`
	Username string       `json:"username"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Role     RoleResponse `json:"role"`
}
