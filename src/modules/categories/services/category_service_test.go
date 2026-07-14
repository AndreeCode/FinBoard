package services

import (
	"context"
	"finboard/src/mocks"
	categoriesDomains "finboard/src/modules/categories/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCategoryService_ObtainCategories_Success(t *testing.T) {
	repo := mocks.NewCategoryRepositoryMock()
	now := time.Now()
	userId := uuid.New()
	catId := uuid.New()
	parentId := uuid.New()
	cat := &categoriesDomains.Category{
		Id:          catId,
		Name:        "Food",
		Description: "Food expenses",
		ParentId:    &parentId,
		UserId:      &userId,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Categories[catId.String()] = cat

	service := NewCategoryService(repo)

	categories, err := service.ObtainCategories(context.Background(), userId.String())

	assert.NoError(t, err)
	assert.Len(t, categories, 1)
	assert.Equal(t, "Food", categories[0].Name)
}

func TestCategoryService_ObtainCategories_Empty(t *testing.T) {
	repo := mocks.NewCategoryRepositoryMock()
	service := NewCategoryService(repo)

	categories, err := service.ObtainCategories(context.Background(), uuid.New().String())

	assert.NoError(t, err)
	assert.Len(t, categories, 0)
}

func TestCategoryService_ObtainCategory_Success(t *testing.T) {
	repo := mocks.NewCategoryRepositoryMock()
	now := time.Now()
	catId := uuid.New()
	cat := &categoriesDomains.Category{
		Id:        catId,
		Name:      "Food",
		UserId:    nil,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.Categories[catId.String()] = cat

	service := NewCategoryService(repo)
	catToFind := &categoriesDomains.Category{Id: catId}

	foundCat, err := service.ObtainCategory(context.Background(), catToFind)

	assert.NoError(t, err)
	assert.NotNil(t, foundCat)
	assert.Equal(t, "Food", foundCat.Name)
}

func TestCategoryService_ObtainCategory_NotFound(t *testing.T) {
	repo := mocks.NewCategoryRepositoryMock()
	service := NewCategoryService(repo)
	catToFind := &categoriesDomains.Category{Id: uuid.New()}

	foundCat, err := service.ObtainCategory(context.Background(), catToFind)

	assert.Error(t, err)
	assert.Nil(t, foundCat)
}

func TestCategoryService_CreateCategory_Success(t *testing.T) {
	repo := mocks.NewCategoryRepositoryMock()
	service := NewCategoryService(repo)
	userId := uuid.New()
	cat := &categoriesDomains.Category{
		Name:     "Food",
		UserId:   &userId,
	}

	createdCat, err := service.CreateCategory(context.Background(), cat)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, createdCat.Id)
	assert.Equal(t, "Food", createdCat.Name)
}

func TestCategoryService_UpdateCategory_Success(t *testing.T) {
	repo := mocks.NewCategoryRepositoryMock()
	now := time.Now()
	catId := uuid.New()
	userId := uuid.New()
	cat := &categoriesDomains.Category{
		Id:        catId,
		Name:      "Food",
		UserId:    &userId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.Categories[catId.String()] = cat

	service := NewCategoryService(repo)
	cat.Name = "Drinks"

	updatedCat, err := service.UpdateCategory(context.Background(), cat)

	assert.NoError(t, err)
	assert.NotNil(t, updatedCat)
	assert.Equal(t, "Drinks", updatedCat.Name)
}

func TestCategoryService_DeleteCategory_Success(t *testing.T) {
	repo := mocks.NewCategoryRepositoryMock()
	now := time.Now()
	catId := uuid.New()
	userId := uuid.New()
	cat := &categoriesDomains.Category{
		Id:        catId,
		Name:      "Food",
		UserId:    &userId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.Categories[catId.String()] = cat

	service := NewCategoryService(repo)

	err := service.DeleteCategory(context.Background(), cat)

	assert.NoError(t, err)
	assert.Len(t, repo.Categories, 0)
}

func TestCategoryService_NewCategoryService(t *testing.T) {
	repo := mocks.NewCategoryRepositoryMock()
	service := NewCategoryService(repo)

	assert.NotNil(t, service)
}
