package apikey

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockApiKeyRepository struct {
	mock.Mock
}

func (m *MockApiKeyRepository) Save(apiKey ApiKey) error {
	args := m.Called(apiKey)

	return args.Error(0)
}

func (m *MockApiKeyRepository) FindById(id string) (ApiKey, error) {
	args := m.Called(id)

	return args.Get(0).(ApiKey), args.Error(1)
}

func (m *MockApiKeyRepository) FindByProjectId(projectId string) (ApiKey, error) {
	args := m.Called(projectId)
	return args.Get(0).(ApiKey), args.Error(1)
}

func (m *MockApiKeyRepository) FindAllByUserIdAndProjectId(userId string, projectId string) ([]ApiKey, error) {
	args := m.Called(userId, projectId)
	return args.Get(0).([]ApiKey), args.Error(1)
}

func (m *MockApiKeyRepository) RevokeApiKey(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockApiKeyRepository) FindProjectIdByKey(apiKey string) (string, error) {
	args := m.Called(apiKey)
	return args.String(0), args.Error(1)
}
func TestUserService_CreateApiKey_Success(t *testing.T) {
	repo := new(MockApiKeyRepository)

	repo.On("Save", mock.Anything).Return(nil)

	service := apiKeyService{
		repo: repo,
	}

	apiKey := &ApiKey{
		Id:   uuid.NewString(),
		Name: "test",
	}

	response, err := service.CreateApiKey(apiKey)

	assert.NoError(t, err)

	assert.Equal(t, "API key created successfully", response.Message)
}

func TestUserService_CreateApiKey_DBError_ReturnsError(t *testing.T) {
	repo := new(MockApiKeyRepository)

	repo.On("Save", mock.Anything).Return(errors.New("DB Failure"))

	service := apiKeyService{
		repo: repo,
	}

	apiKey := &ApiKey{
		Id:   uuid.NewString(),
		Name: "test",
	}

	response, err := service.CreateApiKey(apiKey)

	assert.Error(t, err)

	assert.Empty(t, response)
}
