package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	permissionDomains "finboard/src/modules/permissions/domains"
	permissionRepository "finboard/src/modules/permissions/repository"
)

type PermissionRepositoryIntegrationSuite struct {
	*suite.Suite
	repo *permissionRepository.PermissionRepository
	ctx  context.Context
	*IntegrationTestSuite
}

func TestPermissionRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	integrationSuite, err := SetupPostgresContainer(t)
	require.NoError(t, err)

	s := &PermissionRepositoryIntegrationSuite{
		Suite:                &suite.Suite{},
		IntegrationTestSuite: integrationSuite,
		repo:                 permissionRepository.NewPermissionRepository(),
		ctx:                  integrationSuite.Ctx,
	}

	suite.Run(t, s)
}

func (s *PermissionRepositoryIntegrationSuite) SetupTest() {
	s.IntegrationTestSuite.RunMigrations(s.T())
	s.repo.DB = s.IntegrationTestSuite.DB
}

func (s *PermissionRepositoryIntegrationSuite) TearDownTest() {
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE permissions CASCADE")
}

func (s *PermissionRepositoryIntegrationSuite) TestCreatePermission() {
	permission := &permissionDomains.Permission{
		Id:          uuid.New(),
		Name:        "create_user",
		Description: "Allows creating users",
		CreatedBy:   uuid.New(),
	}

	created, err := s.repo.CreatePermission(s.ctx, permission)

	assert.NoError(s.T(), err)
	assert.NotEqual(s.T(), uuid.Nil, created.Id)
	assert.Equal(s.T(), "create_user", created.Name)
	assert.Equal(s.T(), "Allows creating users", created.Description)
	assert.NotNil(s.T(), created.CreatedAt)
}

func (s *PermissionRepositoryIntegrationSuite) TestCreatePermissionDuplicateName() {
	permission := &permissionDomains.Permission{
		Id:          uuid.New(),
		Name:        "duplicate_test",
		Description: "First permission",
		CreatedBy:   uuid.New(),
	}
	_, err := s.repo.CreatePermission(s.ctx, permission)
	require.NoError(s.T(), err)

	duplicate := &permissionDomains.Permission{
		Id:          uuid.New(),
		Name:        "duplicate_test",
		Description: "Duplicate permission",
		CreatedBy:   uuid.New(),
	}
	_, err = s.repo.CreatePermission(s.ctx, duplicate)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), permissionRepository.ErrPermissionAlreadyExists, err)
}

func (s *PermissionRepositoryIntegrationSuite) TestGetPermission() {
	permission := &permissionDomains.Permission{
		Id:          uuid.New(),
		Name:        "get_test",
		Description: "Get test permission",
		CreatedBy:   uuid.New(),
	}
	created, err := s.repo.CreatePermission(s.ctx, permission)
	require.NoError(s.T(), err)

	found, err := s.repo.GetPermission(s.ctx, &permissionDomains.Permission{Id: created.Id})

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), found)
	assert.Equal(s.T(), created.Id, found.Id)
	assert.Equal(s.T(), "get_test", found.Name)
}

func (s *PermissionRepositoryIntegrationSuite) TestGetPermissionNotFound() {
	nonExistent := &permissionDomains.Permission{Id: uuid.New()}

	found, err := s.repo.GetPermission(s.ctx, nonExistent)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), found)
}

func (s *PermissionRepositoryIntegrationSuite) TestGetListPermissions() {
	for i := 0; i < 3; i++ {
		permission := &permissionDomains.Permission{
			Id:          uuid.New(),
			Name:        "permission_" + string(rune('a'+i)),
			Description: "List test",
			CreatedBy:   uuid.New(),
		}
		_, err := s.repo.CreatePermission(s.ctx, permission)
		require.NoError(s.T(), err)
	}

	permissions, err := s.repo.GetList(s.ctx)

	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(permissions), 3)
}

func (s *PermissionRepositoryIntegrationSuite) TestUpdatePermission() {
	permission := &permissionDomains.Permission{
		Id:          uuid.New(),
		Name:        "before_update",
		Description: "Original description",
		CreatedBy:   uuid.New(),
	}
	created, err := s.repo.CreatePermission(s.ctx, permission)
	require.NoError(s.T(), err)

	newName := "after_update"
	newDescription := "Updated description"
	created.Name = newName
	created.Description = newDescription

	updated, err := s.repo.Update(s.ctx, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), newName, updated.Name)
	assert.Equal(s.T(), newDescription, updated.Description)
}

func (s *PermissionRepositoryIntegrationSuite) TestUpdatePermissionPartial() {
	permission := &permissionDomains.Permission{
		Id:          uuid.New(),
		Name:        "original_name",
		Description: "original_description",
		CreatedBy:   uuid.New(),
	}
	created, err := s.repo.CreatePermission(s.ctx, permission)
	require.NoError(s.T(), err)

	created.Name = "new_name"
	created.Description = "new_description"

	updated, err := s.repo.Update(s.ctx, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "new_name", updated.Name)
	assert.Equal(s.T(), "new_description", updated.Description)
}

func (s *PermissionRepositoryIntegrationSuite) TestPermissionWithLongDescription() {
	longDesc := ""
	for i := 0; i < 500; i++ {
		longDesc += "x"
	}

	permission := &permissionDomains.Permission{
		Id:          uuid.New(),
		Name:        "long_desc_test",
		Description: longDesc,
		CreatedBy:   uuid.New(),
	}

	created, err := s.repo.CreatePermission(s.ctx, permission)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), longDesc, created.Description)
}
