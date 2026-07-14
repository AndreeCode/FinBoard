package domains

type CreateRolePermissionRequest struct {
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}
