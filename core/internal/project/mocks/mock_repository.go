package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/winnerx0/kivia/internal/project"
)


type MockProjectRepository struct {
	mock.Mock
}

func (m *MockProjectRepository) Save(project project.Project) error {
	args := m.Mock.Called(project)
	return args.Error(0)
}

func (m *MockProjectRepository) FindById(id string) (*project.Project, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(*project.Project), args.Error(1)
}

func (m *MockProjectRepository) FindAllByUserId(userId string) ([]project.ProjectResponse, error) {
	args := m.Mock.Called(userId)
	return args.Get(0).([]project.ProjectResponse), args.Error(1)
}

func (m *MockProjectRepository) ExistsById(id string) (bool, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockProjectRepository) FindProjectIdByApiKey(apiKey string) (string, error) {
	args := m.Mock.Called(apiKey)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockProjectRepository) GetAllProjectsByUser(userId string) ([]project.ProjectResponse, error) {
	args := m.Mock.Called(userId)
	return args.Get(0).([]project.ProjectResponse), args.Error(1)
}