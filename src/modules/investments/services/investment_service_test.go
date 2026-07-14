package services

import (
	"context"
	"errors"
	"finboard/src/mocks"
	investmentsDomains "finboard/src/modules/investments/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInvestmentService_ObtainInvestments_Success(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	now := time.Now()
	invId := uuid.New()
	transactionId := uuid.New()
	inv := &investmentsDomains.Investment{
		Id:            invId,
		TransactionId: transactionId,
		ExpectedGain:  15.5,
		RiskLevel:     "medium",
		Status:        "active",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
	repo.Investments[invId.String()] = inv

	service := NewInvestmentService(repo)

	investments, err := service.ObtainInvestments(context.Background(), "")

	assert.NoError(t, err)
	assert.Len(t, investments, 1)
	assert.Equal(t, 15.5, investments[0].ExpectedGain)
}

func TestInvestmentService_ObtainInvestments_Empty(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	service := NewInvestmentService(repo)

	investments, err := service.ObtainInvestments(context.Background(), uuid.New().String())

	assert.NoError(t, err)
	assert.Len(t, investments, 0)
}

func TestInvestmentService_ObtainInvestment_Success(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	now := time.Now()
	invId := uuid.New()
	transactionId := uuid.New()
	inv := &investmentsDomains.Investment{
		Id:            invId,
		TransactionId: transactionId,
		ExpectedGain:  15.5,
		RiskLevel:     "medium",
		Status:        "active",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
	repo.Investments[invId.String()] = inv

	service := NewInvestmentService(repo)
	invToFind := &investmentsDomains.Investment{Id: invId}

	foundInv, err := service.ObtainInvestment(context.Background(), invToFind)

	assert.NoError(t, err)
	assert.NotNil(t, foundInv)
	assert.Equal(t, 15.5, foundInv.ExpectedGain)
}

func TestInvestmentService_ObtainInvestment_NotFound(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	service := NewInvestmentService(repo)
	invToFind := &investmentsDomains.Investment{Id: uuid.New()}

	foundInv, err := service.ObtainInvestment(context.Background(), invToFind)

	assert.Error(t, err)
	assert.Nil(t, foundInv)
}

func TestInvestmentService_CreateInvestment_Success(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	service := NewInvestmentService(repo)
	transactionId := uuid.New()
	inv := &investmentsDomains.Investment{
		TransactionId: transactionId,
		ExpectedGain:  15.5,
		RiskLevel:    "medium",
		Status:       "active",
	}

	createdInv, err := service.CreateInvestment(context.Background(), inv)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, createdInv.Id)
	assert.Equal(t, 15.5, createdInv.ExpectedGain)
}

func TestInvestmentService_UpdateInvestment_Success(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	now := time.Now()
	invId := uuid.New()
	transactionId := uuid.New()
	inv := &investmentsDomains.Investment{
		Id:            invId,
		TransactionId: transactionId,
		ExpectedGain:  15.5,
		RiskLevel:     "medium",
		Status:        "active",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
	repo.Investments[invId.String()] = inv

	service := NewInvestmentService(repo)
	inv.ExpectedGain = 25.0

	updatedInv, err := service.UpdateInvestment(context.Background(), inv)

	assert.NoError(t, err)
	assert.NotNil(t, updatedInv)
	assert.Equal(t, 25.0, updatedInv.ExpectedGain)
}

func TestInvestmentService_UpdateInvestment_NotFound(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	service := NewInvestmentService(repo)
	inv := &investmentsDomains.Investment{
		Id:           uuid.New(),
		ExpectedGain: 25.0,
	}

	updatedInv, err := service.UpdateInvestment(context.Background(), inv)

	assert.Error(t, err)
	assert.Nil(t, updatedInv)
}

func TestInvestmentService_DeleteInvestment_Success(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	now := time.Now()
	invId := uuid.New()
	transactionId := uuid.New()
	inv := &investmentsDomains.Investment{
		Id:            invId,
		TransactionId: transactionId,
		ExpectedGain:  15.5,
		RiskLevel:     "medium",
		Status:        "active",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
	repo.Investments[invId.String()] = inv

	service := NewInvestmentService(repo)

	err := service.DeleteInvestment(context.Background(), inv)

	assert.NoError(t, err)
	assert.Len(t, repo.Investments, 0)
}

func TestInvestmentService_NewInvestmentService(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	service := NewInvestmentService(repo)

	assert.NotNil(t, service)
}

func TestInvestmentService_CheckOwnership_Success(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	service := NewInvestmentService(repo)
	transactionId := uuid.New()
	userId := uuid.New()
	repo.TxUserId = userId

	err := service.CheckOwnership(context.Background(), transactionId, userId.String())

	assert.NoError(t, err)
}

func TestInvestmentService_CheckOwnership_Unauthorized(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	service := NewInvestmentService(repo)
	transactionId := uuid.New()
	userId := uuid.New().String()
	repo.TxUserId = uuid.New()

	err := service.CheckOwnership(context.Background(), transactionId, userId)

	assert.Error(t, err)
	assert.Equal(t, ErrUnauthorized, err)
}

func TestInvestmentService_CheckOwnership_TxUserIdError(t *testing.T) {
	repo := mocks.NewInvestmentRepositoryMock()
	service := NewInvestmentService(repo)
	transactionId := uuid.New()
	userId := uuid.New().String()
	repo.TxUserIdErr = errors.New("database error")

	err := service.CheckOwnership(context.Background(), transactionId, userId)

	assert.Error(t, err)
	assert.NotEqual(t, ErrUnauthorized, err)
}
