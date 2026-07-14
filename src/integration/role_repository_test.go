package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	roleDomains "finboard/src/modules/roles/domains"
	roleRepository "finboard/src/modules/roles/repository"
)

type RoleRepositoryIntegrationSuite struct {
	*suite.Suite
	repo *roleRepository.RoleRepository
	ctx  context.Context
	*IntegrationTestSuite
}

func TestRoleRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	integrationSuite, err := SetupPostgresContainer(t)
	require.NoError(t, err)

	s := &RoleRepositoryIntegrationSuite{
		Suite:                &suite.Suite{},
		IntegrationTestSuite: integrationSuite,
		repo:                 roleRepository.NewRoleRepository(),
		ctx:                  integrationSuite.Ctx,
	}

	suite.Run(t, s)
}

func (s *RoleRepositoryIntegrationSuite) SetupTest() {
	s.IntegrationTestSuite.RunMigrations(s.T())
	s.repo.DB = s.IntegrationTestSuite.DB
}

func (s *RoleRepositoryIntegrationSuite) TearDownTest() {
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE roles CASCADE")
}

func (s *RoleRepositoryIntegrationSuite) TestCreateRole() {
	role := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "admin",
		Description: "Administrator role",
		CreatedBy:   uuid.New(),
	}

	created, err := s.repo.CreateRole(s.ctx, role)

	assert.NoError(s.T(), err)
	assert.NotEqual(s.T(), uuid.Nil, created.Id)
	assert.Equal(s.T(), "admin", created.Name)
	assert.Equal(s.T(), "Administrator role", created.Description)
	assert.NotNil(s.T(), created.CreatedAt)
}

func (s *RoleRepositoryIntegrationSuite) TestCreateRoleDuplicateName() {
	role := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "duplicate_role",
		Description: "First role",
		CreatedBy:   uuid.New(),
	}
	_, err := s.repo.CreateRole(s.ctx, role)
	require.NoError(s.T(), err)

	duplicate := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "duplicate_role",
		Description: "Duplicate role",
		CreatedBy:   uuid.New(),
	}
	_, err = s.repo.CreateRole(s.ctx, duplicate)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), roleRepository.ErrRoleAlreadyExists, err)
}

func (s *RoleRepositoryIntegrationSuite) TestGetRole() {
	role := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "get_test",
		Description: "Get test role",
		CreatedBy:   uuid.New(),
	}
	created, err := s.repo.CreateRole(s.ctx, role)
	require.NoError(s.T(), err)

	found, err := s.repo.GetRole(s.ctx, &roleDomains.Role{Id: created.Id})

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), found)
	assert.Equal(s.T(), created.Id, found.Id)
	assert.Equal(s.T(), "get_test", found.Name)
}

func (s *RoleRepositoryIntegrationSuite) TestGetRoleNotFound() {
	nonExistent := &roleDomains.Role{Id: uuid.New()}

	found, err := s.repo.GetRole(s.ctx, nonExistent)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), found)
}

func (s *RoleRepositoryIntegrationSuite) TestGetRoleByID() {
	role := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "get_by_id_test",
		Description: "Get by ID test",
		CreatedBy:   uuid.New(),
	}
	created, err := s.repo.CreateRole(s.ctx, role)
	require.NoError(s.T(), err)

	found, err := s.repo.GetRoleByID(s.ctx, created.Id.String())

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), found)
	assert.Equal(s.T(), created.Id, found.Id)
	assert.Equal(s.T(), "get_by_id_test", found.Name)
}

func (s *RoleRepositoryIntegrationSuite) TestGetRoleByIDNotFound() {
	nonExistentId := uuid.New().String()

	found, err := s.repo.GetRoleByID(s.ctx, nonExistentId)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), found)
}

func (s *RoleRepositoryIntegrationSuite) TestGetRoleByIDInvalidUUID() {
	invalidId := "not-a-valid-uuid"

	found, err := s.repo.GetRoleByID(s.ctx, invalidId)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), found)
}

func (s *RoleRepositoryIntegrationSuite) TestGetRoleByName() {
	role := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "get_by_name_test",
		Description: "Get by name test",
		CreatedBy:   uuid.New(),
	}
	_, err := s.repo.CreateRole(s.ctx, role)
	require.NoError(s.T(), err)

	found, err := s.repo.GetRoleByName(s.ctx, "get_by_name_test")

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), found)
	assert.Equal(s.T(), "get_by_name_test", found.Name)
}

func (s *RoleRepositoryIntegrationSuite) TestGetRoleByNameNotFound() {
	nonExistentName := "non_existent_role_name"

	found, err := s.repo.GetRoleByName(s.ctx, nonExistentName)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), found)
}

func (s *RoleRepositoryIntegrationSuite) TestGetListRoles() {
	for i := 0; i < 3; i++ {
		role := &roleDomains.Role{
			Id:          uuid.New(),
			Name:        "role_" + string(rune('a'+i)),
			Description: "List test role",
			CreatedBy:   uuid.New(),
		}
		_, err := s.repo.CreateRole(s.ctx, role)
		require.NoError(s.T(), err)
	}

	roles, err := s.repo.GetList(s.ctx)

	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(roles), 3)
}

func (s *RoleRepositoryIntegrationSuite) TestUpdateRole() {
	role := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "before_update",
		Description: "Original description",
		CreatedBy:   uuid.New(),
	}
	created, err := s.repo.CreateRole(s.ctx, role)
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

func (s *RoleRepositoryIntegrationSuite) TestUpdateRolePartial() {
	role := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "original_name",
		Description: "original_description",
		CreatedBy:   uuid.New(),
	}
	created, err := s.repo.CreateRole(s.ctx, role)
	require.NoError(s.T(), err)

	created.Name = "new_name"
	created.Description = "new_description"

	updated, err := s.repo.Update(s.ctx, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "new_name", updated.Name)
	assert.Equal(s.T(), "new_description", updated.Description)
}

func (s *RoleRepositoryIntegrationSuite) TestRoleWithEmptyDescription() {
	role := &roleDomains.Role{
		Id:          uuid.New(),
		Name:        "empty_desc_role",
		Description: "",
		CreatedBy:   uuid.New(),
	}

	created, err := s.repo.CreateRole(s.ctx, role)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "", created.Description)
}
