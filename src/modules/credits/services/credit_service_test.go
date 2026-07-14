package services

import (
	"context"
	"finboard/src/mocks"
	creditsDomains "finboard/src/modules/credits/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreditService_ObtainCredits_Success(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	now := time.Now()
	userId := uuid.New()
	creditId := uuid.New()
	credit := &creditsDomains.Credit{
		Id:           creditId,
		UserId:       userId,
		PersonName:   "John Doe",
		Amount:       1000.50,
		InterestRate: 5.0,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	repo.Credits[creditId.String()] = credit

	service := NewCreditService(repo)

	credits, err := service.ObtainCredits(context.Background(), userId.String())

	assert.NoError(t, err)
	assert.Len(t, credits, 1)
	assert.Equal(t, "John Doe", credits[0].PersonName)
}

func TestCreditService_ObtainCredits_Empty(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	service := NewCreditService(repo)

	credits, err := service.ObtainCredits(context.Background(), uuid.New().String())

	assert.NoError(t, err)
	assert.Len(t, credits, 0)
}

func TestCreditService_ObtainCredit_Success(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	now := time.Now()
	creditId := uuid.New()
	userId := uuid.New()
	credit := &creditsDomains.Credit{
		Id:           creditId,
		UserId:       userId,
		PersonName:   "John Doe",
		Amount:       1000.50,
		InterestRate: 5.0,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	repo.Credits[creditId.String()] = credit

	service := NewCreditService(repo)

	_, err := service.ObtainCredit(context.Background(), creditId)

	assert.NoError(t, err)
}

func TestCreditService_ObtainCredit_NotFound(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	service := NewCreditService(repo)

	_, err := service.ObtainCredit(context.Background(), uuid.New())

	assert.Error(t, err)
}

func TestCreditService_CreateCredit_Success(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	service := NewCreditService(repo)
	userId := uuid.New()
	credit := &creditsDomains.Credit{
		UserId:     userId,
		PersonName: "John Doe",
		Amount:     1000.50,
		Status:     "active",
	}

	createdCredit, err := service.CreateCredit(context.Background(), credit)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, createdCredit.Id)
	assert.Equal(t, "John Doe", createdCredit.PersonName)
}

func TestCreditService_UpdateCredit_Success(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	now := time.Now()
	creditId := uuid.New()
	userId := uuid.New()
	credit := &creditsDomains.Credit{
		Id:           creditId,
		UserId:       userId,
		PersonName:   "John Doe",
		Amount:       1000.50,
		InterestRate: 5.0,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	repo.Credits[creditId.String()] = credit

	service := NewCreditService(repo)
	credit.PersonName = "Jane Doe"

	_, err := service.UpdateCredit(context.Background(), creditId, credit)

	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", credit.PersonName)
}

func TestCreditService_UpdateCredit_NotFound(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	service := NewCreditService(repo)
	credit := &creditsDomains.Credit{
		PersonName: "Jane Doe",
	}

	_, err := service.UpdateCredit(context.Background(), uuid.New(), credit)

	assert.Error(t, err)
}

func TestCreditService_DeleteCredit_Success(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	now := time.Now()
	creditId := uuid.New()
	userId := uuid.New()
	credit := &creditsDomains.Credit{
		Id:           creditId,
		UserId:       userId,
		PersonName:   "John Doe",
		Amount:       1000.50,
		InterestRate: 5.0,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	repo.Credits[creditId.String()] = credit

	service := NewCreditService(repo)

	err := service.DeleteCredit(context.Background(), creditId)

	assert.NoError(t, err)
	assert.Len(t, repo.Credits, 0)
}

func TestCreditService_NewCreditService(t *testing.T) {
	repo := mocks.NewCreditRepositoryMock()
	service := NewCreditService(repo)

	assert.NotNil(t, service)
}
