package repository

import (
	"context"
	"finboard/src/modules/users/domains"
)

func (r *UserRepository) Update(ctx context.Context, domain *domains.User) (*domains.User, error) {
	query := `
		UPDATE users
		SET 
			name = COALESCE($1, name),
			lastname = COALESCE($2, lastname),
			updated_at = COALESCE($3, updated_at)
		WHERE id = $4
	`
	_, err := r.DB.Exec(ctx, query, domain.Name, domain.LastName, domain.UpdatedAt, domain.Id)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
