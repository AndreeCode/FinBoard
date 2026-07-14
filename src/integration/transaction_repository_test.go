package integration

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	transactionDomains "finboard/src/modules/transactions/domains"
	transactionRepository "finboard/src/modules/transactions/repository"
	userDomains "finboard/src/modules/users/domains"
	userRepository "finboard/src/modules/users/repository"
)

type TransactionRepositoryIntegrationSuite struct {
	*suite.Suite
	repo     *transactionRepository.TransactionRepository
	userRepo *userRepository.UserRepository
	ctx      context.Context
	*IntegrationTestSuite
}

func TestTransactionRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	integrationSuite, err := SetupPostgresContainer(t)
	require.NoError(t, err)

	s := &TransactionRepositoryIntegrationSuite{
		Suite:                &suite.Suite{},
		IntegrationTestSuite: integrationSuite,
		repo:                 transactionRepository.NewTransactionRepository(),
		userRepo:             userRepository.NewUserRepository(),
		ctx:                  integrationSuite.Ctx,
	}

	suite.Run(t, s)
}

func (s *TransactionRepositoryIntegrationSuite) SetupTest() {
	s.IntegrationTestSuite.RunMigrations(s.T())
	s.repo.DB = s.IntegrationTestSuite.DB
	s.userRepo.DB = s.IntegrationTestSuite.DB
}

func (s *TransactionRepositoryIntegrationSuite) TearDownTest() {
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE transactions CASCADE")
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE users CASCADE")
}

func (s *TransactionRepositoryIntegrationSuite) createTestUser(t *testing.T) *userDomains.User {
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "TestUser",
		LastName:  "Transaction",
		Email:     "testuser-" + uuid.New().String() + "@example.com",
		Password:  "hashedpassword",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.userRepo.CreateUser(s.ctx, user)
	require.NoError(t, err)
	return user
}

func (s *TransactionRepositoryIntegrationSuite) TestCreateTransaction() {
	user := s.createTestUser(s.T())
	transactionDate := time.Now()

	transaction := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user.Id,
		Amount:          150.50,
		Type:            "income",
		TransactionDate: transactionDate,
		Description:     "Test income",
		CreatedBy:       user.Id,
	}

	created, err := s.repo.CreateTransaction(s.ctx, transaction)

	assert.NoError(s.T(), err)
	assert.NotEqual(s.T(), uuid.Nil, created.Id)
	assert.Equal(s.T(), user.Id, created.UserId)
	assert.Equal(s.T(), 150.50, created.Amount)
	assert.Equal(s.T(), "income", created.Type)
	assert.NotNil(s.T(), created.CreatedAt)
}

func (s *TransactionRepositoryIntegrationSuite) TestCreateTransactionWithCategory() {
	user := s.createTestUser(s.T())
	categoryId := uuid.New()
	transactionDate := time.Now()

	transaction := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user.Id,
		CategoryId:      &categoryId,
		Amount:          200.00,
		Type:            "expense",
		TransactionDate: transactionDate,
		Description:     "Test with category",
		CreatedBy:       user.Id,
	}

	created, err := s.repo.CreateTransaction(s.ctx, transaction)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), created.CategoryId)
	assert.Equal(s.T(), categoryId, *created.CategoryId)
}

func (s *TransactionRepositoryIntegrationSuite) TestGetTransaction() {
	user := s.createTestUser(s.T())
	transactionDate := time.Now()

	transaction := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user.Id,
		Amount:          300.00,
		Type:            "income",
		TransactionDate: transactionDate,
		Description:     "Get test",
		CreatedBy:       user.Id,
	}
	created, err := s.repo.CreateTransaction(s.ctx, transaction)
	require.NoError(s.T(), err)

	found, err := s.repo.GetTransaction(s.ctx, &transactionDomains.Transaction{Id: created.Id})

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), found)
	assert.Equal(s.T(), created.Id, found.Id)
	assert.Equal(s.T(), 300.00, found.Amount)
}

func (s *TransactionRepositoryIntegrationSuite) TestGetTransactionNotFound() {
	nonExistent := &transactionDomains.Transaction{Id: uuid.New()}

	found, err := s.repo.GetTransaction(s.ctx, nonExistent)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), found)
}

func (s *TransactionRepositoryIntegrationSuite) TestGetListTransactions() {
	user := s.createTestUser(s.T())
	transactionDate := time.Now()

	for i := 0; i < 3; i++ {
		transaction := &transactionDomains.Transaction{
			Id:              uuid.New(),
			UserId:          user.Id,
			Amount:          float64(100 + i*50),
			Type:            "income",
			TransactionDate: transactionDate,
			Description:     "List test",
			CreatedBy:       user.Id,
		}
		_, err := s.repo.CreateTransaction(s.ctx, transaction)
		require.NoError(s.T(), err)
	}

	transactions, err := s.repo.GetList(s.ctx, "", "")

	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(transactions), 3)
}

func (s *TransactionRepositoryIntegrationSuite) TestGetListTransactionsByUserId() {
	user1 := s.createTestUser(s.T())
	user2 := s.createTestUser(s.T())
	transactionDate := time.Now()

	tx1 := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user1.Id,
		Amount:          100.00,
		Type:            "income",
		TransactionDate: transactionDate,
		Description:     "User1 tx",
		CreatedBy:       user1.Id,
	}
	_, err := s.repo.CreateTransaction(s.ctx, tx1)
	require.NoError(s.T(), err)

	tx2 := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user2.Id,
		Amount:          200.00,
		Type:            "income",
		TransactionDate: transactionDate,
		Description:     "User2 tx",
		CreatedBy:       user2.Id,
	}
	_, err = s.repo.CreateTransaction(s.ctx, tx2)
	require.NoError(s.T(), err)

	transactions, err := s.repo.GetList(s.ctx, "", user1.Id.String())

	assert.NoError(s.T(), err)
	for _, tx := range transactions {
		assert.Equal(s.T(), user1.Id, tx.UserId)
	}
}

func (s *TransactionRepositoryIntegrationSuite) TestUpdateTransaction() {
	user := s.createTestUser(s.T())
	transactionDate := time.Now()

	transaction := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user.Id,
		Amount:          100.00,
		Type:            "income",
		TransactionDate: transactionDate,
		Description:     "Before update",
		CreatedBy:       user.Id,
	}
	created, err := s.repo.CreateTransaction(s.ctx, transaction)
	require.NoError(s.T(), err)

	newAmount := 500.00
	newDescription := "After update"
	created.Amount = newAmount
	created.Description = newDescription

	updated, err := s.repo.Update(s.ctx, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), newAmount, updated.Amount)
	assert.Equal(s.T(), newDescription, updated.Description)
}

func (s *TransactionRepositoryIntegrationSuite) TestUpdateTransactionPartial() {
	user := s.createTestUser(s.T())
	transactionDate := time.Now()

	transaction := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user.Id,
		Amount:          100.00,
		Type:            "income",
		TransactionDate: transactionDate,
		Description:     "Original description",
		CreatedBy:       user.Id,
	}
	created, err := s.repo.CreateTransaction(s.ctx, transaction)
	require.NoError(s.T(), err)

	originalAmount := created.Amount
	newDescription := "Only description changed"
	created.Description = newDescription

	updated, err := s.repo.Update(s.ctx, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), originalAmount, updated.Amount)
	assert.Equal(s.T(), newDescription, updated.Description)
}

func (s *TransactionRepositoryIntegrationSuite) TestTransactionCanceledFlag() {
	user := s.createTestUser(s.T())
	transactionDate := time.Now()

	transaction := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user.Id,
		Amount:          100.00,
		Type:            "income",
		TransactionDate: transactionDate,
		Canceled:        false,
		Description:     "Canceled test",
		CreatedBy:       user.Id,
	}
	created, err := s.repo.CreateTransaction(s.ctx, transaction)
	require.NoError(s.T(), err)

	created.Canceled = true
	updated, err := s.repo.Update(s.ctx, &created)
	require.NoError(s.T(), err)

	assert.True(s.T(), updated.Canceled)
}

func (s *TransactionRepositoryIntegrationSuite) TestTransactionDates() {
	user := s.createTestUser(s.T())
	transactionDate := time.Now()
	receivedDate := time.Now().Add(24 * time.Hour)
	dueDate := time.Now().Add(48 * time.Hour)

	transaction := &transactionDomains.Transaction{
		Id:              uuid.New(),
		UserId:          user.Id,
		Amount:          100.00,
		Type:            "expense",
		TransactionDate: transactionDate,
		ReceivedDate:    &receivedDate,
		DueDate:         &dueDate,
		Description:     "Dates test",
		CreatedBy:       user.Id,
	}

	created, err := s.repo.CreateTransaction(s.ctx, transaction)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), created.ReceivedDate)
	assert.NotNil(s.T(), created.DueDate)
}
