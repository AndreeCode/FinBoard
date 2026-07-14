package domains

import (
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	Id          uuid.UUID  `json:"id"`
	RoleId      uuid.UUID  `json:"role_id"`
	PermissionId uuid.UUID `json:"permission_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedBy   uuid.UUID  `json:"created_by"`
}
