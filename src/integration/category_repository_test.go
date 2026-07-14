package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	categoryDomains "finboard/src/modules/categories/domains"
	categoryRepository "finboard/src/modules/categories/repository"
	userDomains "finboard/src/modules/users/domains"
	userRepository "finboard/src/modules/users/repository"
)

type CategoryRepositoryIntegrationSuite struct {
	*suite.Suite
	repo     *categoryRepository.CategoryRepository
	userRepo *userRepository.UserRepository
	ctx      context.Context
	*IntegrationTestSuite
}

func TestCategoryRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	integrationSuite, err := SetupPostgresContainer(t)
	require.NoError(t, err)

	s := &CategoryRepositoryIntegrationSuite{
		Suite:                &suite.Suite{},
		IntegrationTestSuite: integrationSuite,
		repo:                 categoryRepository.NewCategoryRepository(),
		userRepo:             userRepository.NewUserRepository(),
		ctx:                  integrationSuite.Ctx,
	}

	suite.Run(t, s)
}

func (s *CategoryRepositoryIntegrationSuite) SetupTest() {
	s.IntegrationTestSuite.RunMigrations(s.T())
	s.repo.DB = s.IntegrationTestSuite.DB
	s.userRepo.DB = s.IntegrationTestSuite.DB
}

func (s *CategoryRepositoryIntegrationSuite) TearDownTest() {
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE categories CASCADE")
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE users CASCADE")
}

func (s *CategoryRepositoryIntegrationSuite) createTestUser(t *testing.T) *userDomains.User {
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "TestUser",
		LastName:  "Category",
		Email:     "testuser-" + uuid.New().String() + "@example.com",
		Password:  "hashedpassword",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.userRepo.CreateUser(s.ctx, user)
	require.NoError(t, err)
	return user
}

func (s *CategoryRepositoryIntegrationSuite) TestCreateCategory() {
	user := s.createTestUser(s.T())

	category := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "TestCategory",
		Description: "Test description",
		UserId:      &user.Id,
		CreatedBy:   user.Id,
	}

	created, err := s.repo.CreateCategory(s.ctx, category)

	assert.NoError(s.T(), err)
	assert.NotEqual(s.T(), uuid.Nil, created.Id)
	assert.Equal(s.T(), "TestCategory", created.Name)
	assert.Equal(s.T(), "Test description", created.Description)
	assert.NotNil(s.T(), created.CreatedAt)
}

func (s *CategoryRepositoryIntegrationSuite) TestCreateGlobalCategory() {
	user := s.createTestUser(s.T())

	category := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "GlobalCategory",
		Description: "Global category without user",
		UserId:      nil,
		CreatedBy:   user.Id,
	}

	created, err := s.repo.CreateCategory(s.ctx, category)

	assert.NoError(s.T(), err)
	assert.Nil(s.T(), created.UserId)
	assert.Equal(s.T(), "GlobalCategory", created.Name)
}

func (s *CategoryRepositoryIntegrationSuite) TestCreateCategoryWithParent() {
	user := s.createTestUser(s.T())

	parentCategory := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "ParentCategory",
		Description: "Parent",
		UserId:      &user.Id,
		CreatedBy:   user.Id,
	}
	parentCreated, err := s.repo.CreateCategory(s.ctx, parentCategory)
	require.NoError(s.T(), err)

	childCategory := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "ChildCategory",
		Description: "Child",
		ParentId:    &parentCreated.Id,
		UserId:      &user.Id,
		CreatedBy:   user.Id,
	}

	childCreated, err := s.repo.CreateCategory(s.ctx, childCategory)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), childCreated.ParentId)
	assert.Equal(s.T(), parentCreated.Id, *childCreated.ParentId)
}

func (s *CategoryRepositoryIntegrationSuite) TestGetCategory() {
	user := s.createTestUser(s.T())

	category := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "GetCategory",
		Description: "Get test",
		UserId:      &user.Id,
		CreatedBy:   user.Id,
	}
	created, err := s.repo.CreateCategory(s.ctx, category)
	require.NoError(s.T(), err)

	found, err := s.repo.GetCategory(s.ctx, &categoryDomains.Category{Id: created.Id})

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), found)
	assert.Equal(s.T(), created.Id, found.Id)
	assert.Equal(s.T(), "GetCategory", found.Name)
}

func (s *CategoryRepositoryIntegrationSuite) TestGetCategoryNotFound() {
	nonExistent := &categoryDomains.Category{Id: uuid.New()}

	found, err := s.repo.GetCategory(s.ctx, nonExistent)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), found)
}

func (s *CategoryRepositoryIntegrationSuite) TestGetListCategories() {
	user := s.createTestUser(s.T())

	for i := 0; i < 3; i++ {
		category := &categoryDomains.Category{
			Id:          uuid.New(),
			Name:        "Category" + string(rune('A'+i)),
			Description: "List test",
			UserId:      &user.Id,
			CreatedBy:   user.Id,
		}
		_, err := s.repo.CreateCategory(s.ctx, category)
		require.NoError(s.T(), err)
	}

	categories, err := s.repo.GetList(s.ctx, "")

	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(categories), 3)
}

func (s *CategoryRepositoryIntegrationSuite) TestGetListCategoriesByUserId() {
	user1 := s.createTestUser(s.T())
	user2 := s.createTestUser(s.T())

	cat1 := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "User1Category",
		Description: "User1",
		UserId:      &user1.Id,
		CreatedBy:   user1.Id,
	}
	_, err := s.repo.CreateCategory(s.ctx, cat1)
	require.NoError(s.T(), err)

	cat2 := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "User2Category",
		Description: "User2",
		UserId:      &user2.Id,
		CreatedBy:   user2.Id,
	}
	_, err = s.repo.CreateCategory(s.ctx, cat2)
	require.NoError(s.T(), err)

	categories, err := s.repo.GetList(s.ctx, user1.Id.String())

	assert.NoError(s.T(), err)
	for _, cat := range categories {
		assert.True(s.T(), cat.UserId == nil || *cat.UserId == user1.Id)
	}
}

func (s *CategoryRepositoryIntegrationSuite) TestUpdateCategory() {
	user := s.createTestUser(s.T())

	category := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "BeforeUpdate",
		Description: "Original description",
		UserId:      &user.Id,
		CreatedBy:   user.Id,
	}
	created, err := s.repo.CreateCategory(s.ctx, category)
	require.NoError(s.T(), err)

	newName := "AfterUpdate"
	newDescription := "New description"
	created.Name = newName
	created.Description = newDescription

	updated, err := s.repo.Update(s.ctx, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), newName, updated.Name)
	assert.Equal(s.T(), newDescription, updated.Description)
}

func (s *CategoryRepositoryIntegrationSuite) TestUpdateCategoryPartial() {
	user := s.createTestUser(s.T())

	category := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "OriginalName",
		Description: "OriginalDescription",
		UserId:      &user.Id,
		CreatedBy:   user.Id,
	}
	created, err := s.repo.CreateCategory(s.ctx, category)
	require.NoError(s.T(), err)

	created.Name = "NewName"
	created.Description = "NewDescription"

	updated, err := s.repo.Update(s.ctx, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "NewName", updated.Name)
	assert.Equal(s.T(), "NewDescription", updated.Description)
}

func (s *CategoryRepositoryIntegrationSuite) TestGlobalCategoryVisibleToAllUsers() {
	user1 := s.createTestUser(s.T())
	user2 := s.createTestUser(s.T())

	globalCategory := &categoryDomains.Category{
		Id:          uuid.New(),
		Name:        "GlobalVisibleCategory",
		Description: "Global category",
		UserId:      nil,
		CreatedBy:   user1.Id,
	}
	_, err := s.repo.CreateCategory(s.ctx, globalCategory)
	require.NoError(s.T(), err)

	categories, err := s.repo.GetList(s.ctx, user2.Id.String())

	assert.NoError(s.T(), err)
	found := false
	for _, cat := range categories {
		if cat.Id == globalCategory.Id {
			found = true
			break
		}
	}
	assert.True(s.T(), found, "Global category should be visible to user2")
}
