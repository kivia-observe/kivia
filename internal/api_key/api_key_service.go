package apikey

import (
	"golang.org/x/crypto/bcrypt"
)

type apiKeyService struct {
	repo Repository
}

func NewApiKeyService(repo Repository) *apiKeyService {
	return &apiKeyService{
		repo: repo,
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
