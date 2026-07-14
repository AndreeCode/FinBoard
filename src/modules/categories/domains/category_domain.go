package domains

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentId    *uuid.UUID `json:"parent_id"`
	UserId      *uuid.UUID `json:"user_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CreatedBy   uuid.UUID  `json:"created_by"`
}
