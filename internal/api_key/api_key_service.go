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

func (s apiKeyService) GetAllApiKeys(userId string) ([]ApiKey, error) {
	return s.repo.FindAllByUserId(userId)
}

func (s apiKeyService) RevokeApiKey(apiKeyId string, userId string) error {
	apiKey, err := s.repo.FindById(apiKeyId)
	if err != nil {
		return err
	}

	if apiKey.Id == "" {
		return ErrApiKeyNotFound
	}

	if apiKey.UserId != userId {
		return ErrUnauthorized
	}

	return s.repo.RevokeApiKey(apiKeyId)
}