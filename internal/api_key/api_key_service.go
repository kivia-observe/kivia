package apikey

import (
	"github.com/gofiber/fiber/v3"
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

func (s apiKeyService) CreateApiKey(c fiber.Ctx, apiKey *ApiKey) error {

	key, err := bcrypt.GenerateFromPassword([]byte(apiKey.Key), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	apiKey.Key = string(key)

	err = s.repo.Save(*apiKey)

	if err != nil {
		return err
	}

	return nil
}
