package services

import (
	"testing"

	dto "github.com/sharpnr-inc/sharpnr-engine/core/dtos"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock Repository Definition ---

// MockOrgRepo is a mock implementation of CoreOrganizationRepository for testing.
type MockOrgRepo struct {
	mock.Mock
}

// We implement the interface methods that our service uses on this mock struct.
func (m *MockOrgRepo) Create(org *models.Organization) error {
	args := m.Called(org)
	return args.Error(0)
}

func (m *MockOrgRepo) FindBySharpnrOrgID(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

// Add mock implementations for other interface methods if the service starts using them.
func (m *MockOrgRepo) Update(org *models.Organization) error {
	args := m.Called(org)
	return args.Error(0)
}

func (m *MockOrgRepo) FindByUUID(uuid string) (*models.Organization, error) {
	args := m.Called(uuid)
	return args.Get(0).(*models.Organization), args.Error(1)
}

func (m *MockOrgRepo) FindByMail(email string) (*models.Organization, error) {
	args := m.Called(email)
	return args.Get(0).(*models.Organization), args.Error(1)
}

// --- The Test Case ---

func TestRegisterOrganization_Success(t *testing.T) {
	// 1. Arrange: Set up the test dependencies.
	mockRepo := new(MockOrgRepo)

	// The service we are testing, with the mock injected.
	service := CreateNewOrganization(mockRepo)

	// The input data for the registration request.
	req := dto.OrganizationRegistrationRequestDTO{
		Name:           "Testaurant",
		RootEmail:      "owner@testaurant.com",
		Password:       "password123",
		MasterPassword: "masterpass",
		// OrgOparationType would be here in a real scenario
	}

	// 2. Set up mock expectations.
	// Expect FindBySharpnrOrgID to be called and return false (ID is unique).
	mockRepo.On("FindBySharpnrOrgID", mock.AnythingOfType("string")).Return(false, nil)

	// Expect Create to be called with any organization model and to succeed.
	mockRepo.On("Create", mock.AnythingOfType("*models.Organization")).Return(nil)

	// 3. Act: Execute the function we want to test.
	resp, token, err := service.RegisterOrganization(req)

	// 4. Assert: Check if the results are what we expect.
	assert.NoError(t, err)    // The error should be nil.
	assert.NotNil(t, resp)    // The response DTO should not be nil.
	assert.NotEmpty(t, token) // The JWT token should not be empty.
	assert.Equal(t, "Testaurant", resp.Name)
	assert.Equal(t, "owner@testaurant.com", resp.RootEmail)

	// Verify that the methods we expected to be called were actually called.
	mockRepo.AssertExpectations(t)
}
