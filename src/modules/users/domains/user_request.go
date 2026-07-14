package domains

type CreateUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   string `json:"role_id"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	LastName *string `json:"last_name"`
}
