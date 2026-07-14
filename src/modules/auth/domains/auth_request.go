package domains

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleId    string `json:"role_id"`
	CreatedBy string `json:"created_by"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthResponse struct {
	User  interface{} `json:"user"`
	Token string      `json:"-"`
}
