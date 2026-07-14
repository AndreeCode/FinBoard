package services

import (
	"context"
	"finboard/src/mocks"
	transactionsDomains "finboard/src/modules/transactions/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_ObtainTransactions_Success(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	now := time.Now()
	userId := uuid.New()
	catId := uuid.New()
	txId := uuid.New()
	tx := &transactionsDomains.Transaction{
		Id:              txId,
		UserId:          userId,
		CategoryId:      &catId,
		Amount:          100.50,
		Type:            "income",
		TransactionDate: now,
		Canceled:        false,
		Description:     "Test transaction",
		CreatedAt:       &now,
		UpdatedAt:       &now,
	}
	repo.Transactions[txId.String()] = tx

	service := NewTransactionService(repo)

	transactions, err := service.ObtainTransactions(context.Background(), "", userId.String())

	assert.NoError(t, err)
	assert.Len(t, transactions, 1)
	assert.Equal(t, 100.50, transactions[0].Amount)
}

func TestTransactionService_ObtainTransactions_Empty(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	service := NewTransactionService(repo)

	transactions, err := service.ObtainTransactions(context.Background(), "", uuid.New().String())

	assert.NoError(t, err)
	assert.Len(t, transactions, 0)
}

func TestTransactionService_ObtainTransaction_Success(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	now := time.Now()
	txId := uuid.New()
	userId := uuid.New()
	catId := uuid.New()
	tx := &transactionsDomains.Transaction{
		Id:              txId,
		UserId:          userId,
		CategoryId:      &catId,
		Amount:          100.50,
		Type:            "income",
		TransactionDate: now,
		Canceled:        false,
		Description:     "Test transaction",
		CreatedAt:       &now,
		UpdatedAt:       &now,
	}
	repo.Transactions[txId.String()] = tx

	service := NewTransactionService(repo)
	txToFind := &transactionsDomains.Transaction{Id: txId}

	foundTx, err := service.ObtainTransaction(context.Background(), txToFind)

	assert.NoError(t, err)
	assert.NotNil(t, foundTx)
	assert.Equal(t, 100.50, foundTx.Amount)
}

func TestTransactionService_ObtainTransaction_NotFound(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	service := NewTransactionService(repo)
	txToFind := &transactionsDomains.Transaction{Id: uuid.New()}

	foundTx, err := service.ObtainTransaction(context.Background(), txToFind)

	assert.Error(t, err)
	assert.Nil(t, foundTx)
}

func TestTransactionService_CreateTransaction_Success(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	service := NewTransactionService(repo)
	userId := uuid.New()
	catId := uuid.New()
	tx := &transactionsDomains.Transaction{
		UserId:          userId,
		CategoryId:      &catId,
		Amount:          100.50,
		Type:            "income",
		TransactionDate: time.Now(),
		Description:     "Test transaction",
	}

	createdTx, err := service.CreateTransaction(context.Background(), tx)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, createdTx.Id)
	assert.Equal(t, 100.50, createdTx.Amount)
}

func TestTransactionService_UpdateTransaction_Success(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	now := time.Now()
	txId := uuid.New()
	userId := uuid.New()
	catId := uuid.New()
	tx := &transactionsDomains.Transaction{
		Id:              txId,
		UserId:          userId,
		CategoryId:      &catId,
		Amount:          100.50,
		Type:            "income",
		TransactionDate: now,
		Canceled:        false,
		Description:     "Test transaction",
		CreatedAt:       &now,
		UpdatedAt:       &now,
	}
	repo.Transactions[txId.String()] = tx

	service := NewTransactionService(repo)
	tx.Amount = 200.00

	updatedTx, err := service.UpdateTransaction(context.Background(), tx)

	assert.NoError(t, err)
	assert.NotNil(t, updatedTx)
	assert.Equal(t, 200.00, updatedTx.Amount)
}

func TestTransactionService_UpdateTransaction_NotFound(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	service := NewTransactionService(repo)
	tx := &transactionsDomains.Transaction{
		Id:     uuid.New(),
		Amount: 200.00,
	}

	updatedTx, err := service.UpdateTransaction(context.Background(), tx)

	assert.Error(t, err)
	assert.Nil(t, updatedTx)
}

func TestTransactionService_DeleteTransaction_Success(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	now := time.Now()
	txId := uuid.New()
	userId := uuid.New()
	catId := uuid.New()
	tx := &transactionsDomains.Transaction{
		Id:              txId,
		UserId:          userId,
		CategoryId:      &catId,
		Amount:          100.50,
		Type:            "income",
		TransactionDate: now,
		Canceled:        false,
		Description:     "Test transaction",
		CreatedAt:       &now,
		UpdatedAt:       &now,
	}
	repo.Transactions[txId.String()] = tx

	service := NewTransactionService(repo)

	err := service.DeleteTransaction(context.Background(), tx)

	assert.NoError(t, err)
	assert.Len(t, repo.Transactions, 0)
}

func TestTransactionService_NewTransactionService(t *testing.T) {
	repo := mocks.NewTransactionRepositoryMock()
	service := NewTransactionService(repo)

	assert.NotNil(t, service)
}
