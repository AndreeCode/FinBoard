package repository

import (
	"context"
	"finboard/src/modules/roles/domains"
	"fmt"
)

func (r *RoleRepository) Update(ctx context.Context, domain *domains.Role) (*domains.Role, error) {
	query := `
		UPDATE roles
		SET 
			name = COALESCE($1, name),
			description = COALESCE($2, description),
			updated_at = COALESCE($3, updated_at)
		WHERE id = $4
	`
	_, err := r.DB.Exec(ctx, query, domain.Name, domain.Description, domain.UpdatedAt, domain.Id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err

	}

	return domain, nil
}
