package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	userDomains "finboard/src/modules/users/domains"
	userRepository "finboard/src/modules/users/repository"
)

type UserRepositoryIntegrationSuite struct {
	*suite.Suite
	repo *userRepository.UserRepository
	ctx  context.Context
	*IntegrationTestSuite
}

func TestUserRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	integrationSuite, err := SetupPostgresContainer(t)
	require.NoError(t, err)

	s := &UserRepositoryIntegrationSuite{
		Suite:                &suite.Suite{},
		IntegrationTestSuite: integrationSuite,
		repo:                 userRepository.NewUserRepository(),
		ctx:                  integrationSuite.Ctx,
	}

	suite.Run(t, s)
}

func (s *UserRepositoryIntegrationSuite) SetupTest() {
	s.IntegrationTestSuite.RunMigrations(s.T())
	s.repo.DB = s.IntegrationTestSuite.DB
}

func (s *UserRepositoryIntegrationSuite) TearDownTest() {
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE users CASCADE")
}

func (s *UserRepositoryIntegrationSuite) TestCreateUser() {
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "John",
		LastName:  "Doe",
		Email:     "john-" + uuid.New().String() + "@example.com",
		Password:  "hashedpassword",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}

	createdUser, err := s.repo.CreateUser(s.ctx, user)

	assert.NoError(s.T(), err)
	assert.NotEqual(s.T(), uuid.Nil, createdUser.Id)
	assert.Equal(s.T(), "John", createdUser.Name)
	assert.NotNil(s.T(), createdUser.CreatedAt)
}

func (s *UserRepositoryIntegrationSuite) TestGetUserByEmail() {
	email := "test-" + uuid.New().String() + "@example.com"
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "Jane",
		LastName:  "Smith",
		Email:     email,
		Password:  "hashedpassword",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.repo.CreateUser(s.ctx, user)
	assert.NoError(s.T(), err)

	foundUser, err := s.repo.GetUserByEmail(s.ctx, email)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), foundUser)
	assert.Equal(s.T(), email, foundUser.Email)
	assert.Equal(s.T(), "Jane", foundUser.Name)
}

func (s *UserRepositoryIntegrationSuite) TestGetUsers() {
	users := []userDomains.User{
		{
			Id:        uuid.New(),
			Name:      "User1",
			LastName:  "Test",
			Email:     "user1-" + uuid.New().String() + "@example.com",
			Password:  "pass",
			RoleId:    uuid.New(),
			CreatedBy: uuid.New(),
		},
		{
			Id:        uuid.New(),
			Name:      "User2",
			LastName:  "Test",
			Email:     "user2-" + uuid.New().String() + "@example.com",
			Password:  "pass",
			RoleId:    uuid.New(),
			CreatedBy: uuid.New(),
		},
	}

	for _, u := range users {
		_, err := s.repo.CreateUser(s.ctx, &u)
		assert.NoError(s.T(), err)
	}

	allUsers, err := s.repo.GetList(s.ctx)

	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(allUsers), 2)
}

func (s *UserRepositoryIntegrationSuite) TestUpdateLastLogin() {
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "LoginTest",
		LastName:  "User",
		Email:     "logintest-" + uuid.New().String() + "@example.com",
		Password:  "pass",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.repo.CreateUser(s.ctx, user)
	assert.NoError(s.T(), err)
	user.LastLogin = nil

	err = s.repo.UpdateLastLogin(s.ctx, user)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), user.LastLogin)
}

func (s *UserRepositoryIntegrationSuite) TestGetUserByID() {
	email := "getbyid-" + uuid.New().String() + "@example.com"
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "GetByID",
		LastName:  "Test",
		Email:     email,
		Password:  "hashedpassword",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	createdUser, err := s.repo.CreateUser(s.ctx, user)
	assert.NoError(s.T(), err)

	foundUser, err := s.repo.GetUserByID(s.ctx, createdUser.Id.String())

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), foundUser)
	assert.Equal(s.T(), createdUser.Id, foundUser.Id)
	assert.Equal(s.T(), email, foundUser.Email)
	assert.Equal(s.T(), "GetByID", foundUser.Name)
	assert.Equal(s.T(), "Test", foundUser.LastName)
}

func (s *UserRepositoryIntegrationSuite) TestGetUserByIDNotFound() {
	nonExistentID := uuid.New().String()

	foundUser, err := s.repo.GetUserByID(s.ctx, nonExistentID)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), foundUser)
}

func (s *UserRepositoryIntegrationSuite) TestGetUserByIDInvalidUUID() {
	invalidID := "not-a-valid-uuid"

	foundUser, err := s.repo.GetUserByID(s.ctx, invalidID)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), foundUser)
}

func (s *UserRepositoryIntegrationSuite) TestGetUser() {
	email := "getuser-" + uuid.New().String() + "@example.com"
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "GetUser",
		LastName:  "Test",
		Email:     email,
		Password:  "hashedpassword",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	createdUser, err := s.repo.CreateUser(s.ctx, user)
	assert.NoError(s.T(), err)

	foundUser, err := s.repo.GetUser(s.ctx, &userDomains.User{Id: createdUser.Id})

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), foundUser)
	assert.Equal(s.T(), createdUser.Id, foundUser.Id)
	assert.Equal(s.T(), email, foundUser.Email)
}

func (s *UserRepositoryIntegrationSuite) TestGetUserNotFound() {
	nonExistentUser := &userDomains.User{Id: uuid.New()}

	foundUser, err := s.repo.GetUser(s.ctx, nonExistentUser)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), foundUser)
}

func (s *UserRepositoryIntegrationSuite) TestUpdate() {
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "BeforeUpdate",
		LastName:  "User",
		Email:     "update-" + uuid.New().String() + "@example.com",
		Password:  "pass",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.repo.CreateUser(s.ctx, user)
	assert.NoError(s.T(), err)

	newName := "AfterUpdate"
	newLastName := "UpdatedLastName"
	user.Name = newName
	user.LastName = newLastName

	updatedUser, err := s.repo.Update(s.ctx, user)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), updatedUser)
	assert.Equal(s.T(), newName, updatedUser.Name)
	assert.Equal(s.T(), newLastName, updatedUser.LastName)

	reloadedUser, err := s.repo.GetUserByID(s.ctx, user.Id.String())
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), newName, reloadedUser.Name)
	assert.Equal(s.T(), newLastName, reloadedUser.LastName)
}

func (s *UserRepositoryIntegrationSuite) TestUpdatePartialNameOnly() {
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "OriginalName",
		LastName:  "OriginalLastName",
		Email:     "updatepartial-" + uuid.New().String() + "@example.com",
		Password:  "pass",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.repo.CreateUser(s.ctx, user)
	assert.NoError(s.T(), err)

	user.Name = "OnlyNameChanged"
	originalLastName := user.LastName

	_, err = s.repo.Update(s.ctx, user)

	assert.NoError(s.T(), err)

	reloadedUser, err := s.repo.GetUserByID(s.ctx, user.Id.String())
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "OnlyNameChanged", reloadedUser.Name)
	assert.Equal(s.T(), originalLastName, reloadedUser.LastName)
}

func (s *UserRepositoryIntegrationSuite) TestDuplicateEmail() {
	email := "duplicate-" + uuid.New().String() + "@example.com"
	user1 := &userDomains.User{
		Id:        uuid.New(),
		Name:      "User1",
		LastName:  "Test",
		Email:     email,
		Password:  "pass",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.repo.CreateUser(s.ctx, user1)
	assert.NoError(s.T(), err)

	user2 := &userDomains.User{
		Id:        uuid.New(),
		Name:      "User2",
		LastName:  "Test",
		Email:     email,
		Password:  "pass",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err = s.repo.CreateUser(s.ctx, user2)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), userRepository.ErrUserAlreadyExists, err)
}

func (s *UserRepositoryIntegrationSuite) TestGetUserByEmailNotFound() {
	nonExistentEmail := "nonexistent-" + uuid.New().String() + "@example.com"

	foundUser, err := s.repo.GetUserByEmail(s.ctx, nonExistentEmail)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), foundUser)
}

func (s *UserRepositoryIntegrationSuite) TestGetUserByEmailRaw() {
	email := "raw-" + uuid.New().String() + "@example.com"
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "RawTest",
		LastName:  "User",
		Email:     email,
		Password:  "pass",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.repo.CreateUser(s.ctx, user)
	assert.NoError(s.T(), err)

	foundUser, err := s.repo.GetUserByEmailRaw(s.ctx, email)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), foundUser)
	assert.Equal(s.T(), email, foundUser.Email)
}

func (s *UserRepositoryIntegrationSuite) TestGetUserByEmailRawNotFound() {
	nonExistentEmail := "rawnotfound-" + uuid.New().String() + "@example.com"

	foundUser, err := s.repo.GetUserByEmailRaw(s.ctx, nonExistentEmail)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), foundUser)
}

func (s *UserRepositoryIntegrationSuite) TestGetListMultipleUsers() {
	users := []userDomains.User{
		{
			Id:        uuid.New(),
			Name:      "Multi1",
			LastName:  "User",
			Email:     "multi1-" + uuid.New().String() + "@example.com",
			Password:  "pass",
			RoleId:    uuid.New(),
			CreatedBy: uuid.New(),
		},
		{
			Id:        uuid.New(),
			Name:      "Multi2",
			LastName:  "User",
			Email:     "multi2-" + uuid.New().String() + "@example.com",
			Password:  "pass",
			RoleId:    uuid.New(),
			CreatedBy: uuid.New(),
		},
		{
			Id:        uuid.New(),
			Name:      "Multi3",
			LastName:  "User",
			Email:     "multi3-" + uuid.New().String() + "@example.com",
			Password:  "pass",
			RoleId:    uuid.New(),
			CreatedBy: uuid.New(),
		},
	}

	for i := range users {
		_, err := s.repo.CreateUser(s.ctx, &users[i])
		assert.NoError(s.T(), err)
	}

	allUsers, err := s.repo.GetList(s.ctx)

	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(allUsers), 3)
}

func (s *UserRepositoryIntegrationSuite) TestCreateUserReturnsAllFields() {
	email := "allfields-" + uuid.New().String() + "@example.com"
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "AllFields",
		LastName:  "Test",
		Email:     email,
		Password:  "hashedpassword123",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}

	createdUser, err := s.repo.CreateUser(s.ctx, user)

	assert.NoError(s.T(), err)
	assert.NotEqual(s.T(), uuid.Nil, createdUser.Id)
	assert.Equal(s.T(), "AllFields", createdUser.Name)
	assert.Equal(s.T(), "Test", createdUser.LastName)
	assert.Equal(s.T(), email, createdUser.Email)
	assert.Equal(s.T(), "hashedpassword123", createdUser.Password)
	assert.NotNil(s.T(), createdUser.RoleId)
	assert.NotNil(s.T(), createdUser.CreatedAt)
	assert.NotNil(s.T(), createdUser.UpdatedAt)
	assert.Nil(s.T(), createdUser.DeletedAt)
	assert.Nil(s.T(), createdUser.LastLogin)
}

func (s *UserRepositoryIntegrationSuite) TestUpdateLastLoginPersists() {
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "LoginPersist",
		LastName:  "User",
		Email:     "loginpersist-" + uuid.New().String() + "@example.com",
		Password:  "pass",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.repo.CreateUser(s.ctx, user)
	assert.NoError(s.T(), err)

	err = s.repo.UpdateLastLogin(s.ctx, user)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), user.LastLogin)

	reloadedUser, err := s.repo.GetUserByID(s.ctx, user.Id.String())
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), reloadedUser.LastLogin)
}
