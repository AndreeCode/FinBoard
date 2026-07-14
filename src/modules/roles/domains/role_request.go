package domains

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type UpdateRoleRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
