package apikey

import (
	"log"

	"github.com/winnerx0/dyno/internal/project"
	"github.com/winnerx0/dyno/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type apiKeyService struct {
	repo Repository
	projectRepo project.Repository
}

func NewApiKeyService(repo Repository, projectRepo project.Repository) *apiKeyService {
	return &apiKeyService{
		repo: repo,
		projectRepo: projectRepo,
	}
}

func (s apiKeyService) CreateApiKey(apiKey *ApiKey) (createApiKeyResponse, error) {

	key, err := generateAPIKey()

	if err != nil {
		return createApiKeyResponse{}, err
	}

	hashedKey, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		return createApiKeyResponse{}, err
	}

	apiKey.Key = string(hashedKey)

	err = s.repo.Save(*apiKey)

	if err != nil {
		return createApiKeyResponse{}, err
	}

	return createApiKeyResponse{
		Message: "API key created successfully",
		ApiKey: key,
	}, nil
}

func (s apiKeyService) GetAllApiKeysByProject(userId string, projectId string) ([]ApiKey, error) {

	projectExists, err := s.projectRepo.ExistsById(projectId)

	if err != nil {

		log.Println("Error: Error find project using id", err.Error())
		return nil, utils.ErrProjectNotFound
	}


	if !projectExists {
		return nil, utils.ErrProjectNotFound
	}

	return s.repo.FindAllByUserIdAndProjectId(userId, projectId)
}

func (s apiKeyService) RevokeApiKey(apiKeyId string, userId string) error {
	apiKey, err := s.repo.FindById(apiKeyId)
	if err != nil {
		return err
	}
	
	if apiKey.Id == "" {
		return utils.ErrApiKeyNotFound
	}

	if apiKey.UserId != userId {
		return utils.ErrUnauthorized
	}

	return s.repo.RevokeApiKey(apiKeyId)
}