package middleware

import (

	"github.com/gofiber/fiber/v3"
	apikey "github.com/winnerx0/dyno/internal/api_key"
	"golang.org/x/crypto/bcrypt"
)

type apiKeyMiddleware struct {
	repo apikey.Repository
}

func NewAuthMiddleware(repo apikey.Repository) *apiKeyMiddleware{
	return &apiKeyMiddleware{
		repo: repo,
	}
}
 
func (m apiKeyMiddleware) APiKeyMiddleware(c fiber.Ctx) error {

	apiKey := c.Get("x-dyno-api-key")

	if apiKey == "" {
		return c.Status(403).JSON("Invaild API Key")
	}

	hashedKey := hashApiKey(apiKey)

	key, err := m.repo.FindByKey(hashedKey)

	if err != nil {
		return c.Status(403).JSON("Invaild API Key")
	}

	if key.Id == "" {
		return c.Status(403).JSON("Invaild API Key")
	}

	return c.Next()

}

func hashApiKey(apiKey string) string {
	
	hashedKey, err := bcrypt.GenerateFromPassword([]byte(apiKey), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedKey)
}