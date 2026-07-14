package integration

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	creditDomains "finboard/src/modules/credits/domains"
	creditRepository "finboard/src/modules/credits/repository"
	userDomains "finboard/src/modules/users/domains"
	userRepository "finboard/src/modules/users/repository"
)

type CreditRepositoryIntegrationSuite struct {
	*suite.Suite
	repo     *creditRepository.CreditRepository
	userRepo *userRepository.UserRepository
	ctx      context.Context
	*IntegrationTestSuite
}

func TestCreditRepositoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	integrationSuite, err := SetupPostgresContainer(t)
	require.NoError(t, err)

	s := &CreditRepositoryIntegrationSuite{
		Suite:                &suite.Suite{},
		IntegrationTestSuite: integrationSuite,
		repo:                 creditRepository.NewCreditRepository(),
		userRepo:             userRepository.NewUserRepository(),
		ctx:                  integrationSuite.Ctx,
	}

	suite.Run(t, s)
}

func (s *CreditRepositoryIntegrationSuite) SetupTest() {
	s.IntegrationTestSuite.RunMigrations(s.T())
	s.repo.DB = s.IntegrationTestSuite.DB
	s.userRepo.DB = s.IntegrationTestSuite.DB
}

func (s *CreditRepositoryIntegrationSuite) TearDownTest() {
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE credits CASCADE")
	s.IntegrationTestSuite.DB.Exec(s.Ctx, "TRUNCATE TABLE users CASCADE")
}

func (s *CreditRepositoryIntegrationSuite) createTestUser(t *testing.T) *userDomains.User {
	user := &userDomains.User{
		Id:        uuid.New(),
		Name:      "TestUser",
		LastName:  "Credit",
		Email:     "testuser-" + uuid.New().String() + "@example.com",
		Password:  "hashedpassword",
		RoleId:    uuid.New(),
		CreatedBy: uuid.New(),
	}
	_, err := s.userRepo.CreateUser(s.ctx, user)
	require.NoError(t, err)
	return user
}

func (s *CreditRepositoryIntegrationSuite) TestCreateCredit() {
	user := s.createTestUser(s.T())

	credit := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user.Id,
		PersonName:   "John Doe",
		Amount:       1000.50,
		InterestRate: 5.25,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedBy:    user.Id,
	}

	created, err := s.repo.CreateCredit(s.ctx, credit)

	assert.NoError(s.T(), err)
	assert.NotEqual(s.T(), uuid.Nil, created.Id)
	assert.Equal(s.T(), "John Doe", created.PersonName)
	assert.Equal(s.T(), 1000.50, created.Amount)
	assert.Equal(s.T(), 5.25, created.InterestRate)
	assert.True(s.T(), created.IsCreditor)
	assert.False(s.T(), created.IsSecure)
	assert.Equal(s.T(), "active", created.Status)
	assert.NotNil(s.T(), created.CreatedAt)
}

func (s *CreditRepositoryIntegrationSuite) TestCreateCreditWithDueDate() {
	user := s.createTestUser(s.T())
	dueDate := time.Now().Add(30 * 24 * time.Hour)

	credit := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user.Id,
		PersonName:   "Jane Doe",
		Amount:       2000.00,
		InterestRate: 10.00,
		IsCreditor:   false,
		IsSecure:     true,
		DueDate:      &dueDate,
		Status:       "active",
		CreatedBy:    user.Id,
	}

	created, err := s.repo.CreateCredit(s.ctx, credit)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), created.DueDate)
	assert.True(s.T(), created.IsSecure)
	assert.False(s.T(), created.IsCreditor)
}

func (s *CreditRepositoryIntegrationSuite) TestObtainCredit() {
	user := s.createTestUser(s.T())

	credit := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user.Id,
		PersonName:   "Obtain Test",
		Amount:       500.00,
		InterestRate: 3.00,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedBy:    user.Id,
	}
	created, err := s.repo.CreateCredit(s.ctx, credit)
	require.NoError(s.T(), err)

	obtained, err := s.repo.ObtainCredit(s.ctx, created.Id)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), created.Id, obtained.Id)
	assert.Equal(s.T(), "Obtain Test", obtained.PersonName)
	assert.Equal(s.T(), 500.00, obtained.Amount)
}

func (s *CreditRepositoryIntegrationSuite) TestObtainCreditNotFound() {
	nonExistentId := uuid.New()

	_, err := s.repo.ObtainCredit(s.ctx, nonExistentId)

	assert.Error(s.T(), err)
}

func (s *CreditRepositoryIntegrationSuite) TestObtainCredits() {
	user := s.createTestUser(s.T())

	for i := 0; i < 3; i++ {
		credit := &creditDomains.Credit{
			Id:           uuid.New(),
			UserId:       user.Id,
			PersonName:   "Person" + string(rune('A'+i)),
			Amount:       float64(100 * (i + 1)),
			InterestRate: float64(i) + 1.0,
			IsCreditor:   i%2 == 0,
			IsSecure:     false,
			Status:       "active",
			CreatedBy:    user.Id,
		}
		_, err := s.repo.CreateCredit(s.ctx, credit)
		require.NoError(s.T(), err)
	}

	credits, err := s.repo.ObtainCredits(s.ctx, "")

	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(credits), 3)
}

func (s *CreditRepositoryIntegrationSuite) TestObtainCreditsByUserId() {
	user1 := s.createTestUser(s.T())
	user2 := s.createTestUser(s.T())

	credit1 := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user1.Id,
		PersonName:   "User1Person",
		Amount:       100.00,
		InterestRate: 1.00,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedBy:    user1.Id,
	}
	_, err := s.repo.CreateCredit(s.ctx, credit1)
	require.NoError(s.T(), err)

	credit2 := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user2.Id,
		PersonName:   "User2Person",
		Amount:       200.00,
		InterestRate: 2.00,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedBy:    user2.Id,
	}
	_, err = s.repo.CreateCredit(s.ctx, credit2)
	require.NoError(s.T(), err)

	credits, err := s.repo.ObtainCredits(s.ctx, user1.Id.String())

	assert.NoError(s.T(), err)
	for _, c := range credits {
		assert.Equal(s.T(), user1.Id, c.UserId)
	}
}

func (s *CreditRepositoryIntegrationSuite) TestUpdateCredit() {
	user := s.createTestUser(s.T())

	credit := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user.Id,
		PersonName:   "BeforeUpdate",
		Amount:       100.00,
		InterestRate: 1.00,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedBy:    user.Id,
	}
	created, err := s.repo.CreateCredit(s.ctx, credit)
	require.NoError(s.T(), err)

	newPersonName := "AfterUpdate"
	newAmount := 500.00
	newStatus := "paid"
	created.PersonName = newPersonName
	created.Amount = newAmount
	created.Status = newStatus

	updated, err := s.repo.UpdateCredit(s.ctx, created.Id, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), newPersonName, updated.PersonName)
	assert.Equal(s.T(), newAmount, updated.Amount)
	assert.Equal(s.T(), newStatus, updated.Status)
}

func (s *CreditRepositoryIntegrationSuite) TestUpdateCreditPartial() {
	user := s.createTestUser(s.T())

	credit := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user.Id,
		PersonName:   "OriginalPerson",
		Amount:       100.00,
		InterestRate: 1.00,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedBy:    user.Id,
	}
	created, err := s.repo.CreateCredit(s.ctx, credit)
	require.NoError(s.T(), err)

	originalAmount := created.Amount
	originalStatus := created.Status
	created.PersonName = "NewPersonName"

	updated, err := s.repo.UpdateCredit(s.ctx, created.Id, &created)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "NewPersonName", updated.PersonName)
	assert.Equal(s.T(), originalAmount, updated.Amount)
	assert.Equal(s.T(), originalStatus, updated.Status)
}

func (s *CreditRepositoryIntegrationSuite) TestCreditStatusTransitions() {
	user := s.createTestUser(s.T())

	credit := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user.Id,
		PersonName:   "StatusTest",
		Amount:       1000.00,
		InterestRate: 5.00,
		IsCreditor:   true,
		IsSecure:     false,
		Status:       "active",
		CreatedBy:    user.Id,
	}
	created, err := s.repo.CreateCredit(s.ctx, credit)
	require.NoError(s.T(), err)

	statuses := []string{"active", "paid", "overdue", "cancelled"}
	for _, status := range statuses {
		created.Status = status
		updated, err := s.repo.UpdateCredit(s.ctx, created.Id, &created)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), status, updated.Status)
	}
}

func (s *CreditRepositoryIntegrationSuite) TestCreditInterestRate() {
	user := s.createTestUser(s.T())

	credit := &creditDomains.Credit{
		Id:           uuid.New(),
		UserId:       user.Id,
		PersonName:   "InterestTest",
		Amount:       1000.00,
		InterestRate: 15.75,
		IsCreditor:   false,
		IsSecure:     true,
		Status:       "active",
		CreatedBy:    user.Id,
	}

	created, err := s.repo.CreateCredit(s.ctx, credit)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 15.75, created.InterestRate)
	assert.False(s.T(), created.IsCreditor)
	assert.True(s.T(), created.IsSecure)
}
