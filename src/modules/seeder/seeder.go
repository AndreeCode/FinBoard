package seeder

import (
	"context"
	"errors"
	"finboard/src/core/utils"
	permDomain "finboard/src/modules/permissions/domains"
	permRepo "finboard/src/modules/permissions/repository"
	rolePermDomain "finboard/src/modules/role_permissions/domains"
	rolePermRepo "finboard/src/modules/role_permissions/repository"
	roleDomain "finboard/src/modules/roles/domains"
	roleRepo "finboard/src/modules/roles/repository"
	userDomain "finboard/src/modules/users/domains"
	userRepo "finboard/src/modules/users/repository"

	"github.com/google/uuid"
)

type Seeder struct {
	roleRepo     *roleRepo.RoleRepository
	permRepo     *permRepo.PermissionRepository
	rolePermRepo *rolePermRepo.RolePermissionRepository
	userRepo     *userRepo.UserRepository
}

func NewSeeder() *Seeder {
	return &Seeder{
		roleRepo:     roleRepo.NewRoleRepository(),
		permRepo:     permRepo.NewPermissionRepository(),
		rolePermRepo: rolePermRepo.NewRolePermissionRepository(),
		userRepo:     userRepo.NewUserRepository(),
	}
}

func (s *Seeder) Seed(ctx context.Context) error {
	adminRoleID, publicRoleID, err := s.seedRoles(ctx)
	if err != nil {
		return err
	}

	permIDs, err := s.seedPermissions(ctx)
	if err != nil {
		return err
	}

	err = s.seedRolePermissions(ctx, adminRoleID, permIDs)
	if err != nil {
		return err
	}

	err = s.seedPublicRolePermissions(ctx, publicRoleID)
	if err != nil {
		return err
	}

	_, err = s.seedAdminUser(ctx, adminRoleID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedRoles(ctx context.Context) (uuid.UUID, uuid.UUID, error) {
	adminRoleID, err := s.getOrCreateRole(ctx, "admin", "Administrator role with all permissions")
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	publicRoleID, err := s.getOrCreateRole(ctx, "public", "Public role for unregistered users")
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return adminRoleID, publicRoleID, nil
}

func (s *Seeder) seedPublicRolePermissions(ctx context.Context, publicRoleID uuid.UUID) error {
	nilUUID := uuid.Nil

	publicPerms := []string{"read_users", "update_users"}

	for _, permName := range publicPerms {
		perm, err := s.permRepo.GetPermissionByName(ctx, permName)
		if err != nil {
			continue
		}

		rp := &rolePermDomain.RolePermission{
			RoleId:       publicRoleID,
			PermissionId: perm.Id,
			CreatedBy:    nilUUID,
		}

		_, err = s.rolePermRepo.CreateRolePermission(ctx, rp)
		if err != nil {
			continue
		}
	}

	return nil
}

func (s *Seeder) getOrCreateRole(ctx context.Context, name, description string) (uuid.UUID, error) {
	existingRole, err := s.roleRepo.GetRoleByName(ctx, name)
	if err == nil && existingRole != nil {
		return existingRole.Id, nil
	}

	nilUUID := uuid.Nil
	role := &roleDomain.Role{
		Name:        name,
		Description: description,
		CreatedBy:   nilUUID,
	}

	roleResult, err := s.roleRepo.CreateRole(ctx, role)
	if err != nil {
		if errors.Is(err, roleRepo.ErrRoleAlreadyExists) {
			existingRole, err := s.roleRepo.GetRoleByName(ctx, name)
			if err != nil {
				return uuid.Nil, err
			}
			return existingRole.Id, nil
		}
		return uuid.Nil, err
	}

	return roleResult.Id, nil
}

func (s *Seeder) seedPermissions(ctx context.Context) ([]uuid.UUID, error) {
	modules := []string{"roles", "users", "categories", "transactions", "investments", "credits", "permissions", "role_permissions"}
	actions := []string{"create", "read", "update", "delete"}

	var permIDs []uuid.UUID
	nilUUID := uuid.Nil

	for _, module := range modules {
		for _, action := range actions {
			permName := action + "_" + module

			existingPerm, err := s.permRepo.GetPermissionByName(ctx, permName)
			if err == nil && existingPerm != nil {
				permIDs = append(permIDs, existingPerm.Id)
				continue
			}

			perm := &permDomain.Permission{
				Name:        permName,
				Description: "Permission to " + action + " " + module,
				CreatedBy:   nilUUID,
			}

			p, err := s.permRepo.CreatePermission(ctx, perm)
			if err != nil {
				if errors.Is(err, permRepo.ErrPermissionAlreadyExists) {
					existingPerm, err := s.permRepo.GetPermissionByName(ctx, permName)
					if err != nil {
						return nil, err
					}
					permIDs = append(permIDs, existingPerm.Id)
					continue
				}
				return nil, err
			}
			permIDs = append(permIDs, p.Id)
		}
	}

	return permIDs, nil
}

func (s *Seeder) seedRolePermissions(ctx context.Context, roleID uuid.UUID, permIDs []uuid.UUID) error {
	nilUUID := uuid.Nil

	for _, permID := range permIDs {
		rp := &rolePermDomain.RolePermission{
			RoleId:       roleID,
			PermissionId: permID,
			CreatedBy:   nilUUID,
		}

		_, err := s.rolePermRepo.CreateRolePermission(ctx, rp)
		if err != nil {
			continue
		}
	}

	return nil
}

func (s *Seeder) seedAdminUser(ctx context.Context, roleID uuid.UUID) (uuid.UUID, error) {
	existingUser, err := s.userRepo.GetUserByEmailRaw(ctx, "admin@finboard.com")
	if err == nil && existingUser != nil {
		return existingUser.Id, nil
	}

	hashedPassword, err := utils.HashPassword("admin123", nil)
	if err != nil {
		return uuid.Nil, err
	}

	nilUUID := uuid.Nil
	adminUser := &userDomain.User{
		Name:      "Admin",
		LastName:  "User",
		Email:     "admin@finboard.com",
		Password:  hashedPassword,
		RoleId:    roleID,
		CreatedBy: nilUUID,
	}

	adminUserResult, err := s.userRepo.CreateUser(ctx, adminUser)
	if err != nil {
		if errors.Is(err, userRepo.ErrUserAlreadyExists) {
			existingUser, err := s.userRepo.GetUserByEmailRaw(ctx, "admin@finboard.com")
			if err != nil {
				return uuid.Nil, err
			}
			return existingUser.Id, nil
		}
		return uuid.Nil, err
	}

	return adminUserResult.Id, nil
}
