package domains

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	RoleId    uuid.UUID  `json:"role_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedBy uuid.UUID  `json:"created_by"`
	LastLogin *time.Time `json:"last_login"`
}
