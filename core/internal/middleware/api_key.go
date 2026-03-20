package middleware

import (

	"github.com/gofiber/fiber/v3"
	apikey "github.com/winnerx0/dyno/internal/api_key"
)

type apiKeyMiddleware struct {
	repo *apikey.Repository
}

func NewApiKeyMiddleware(repo *apikey.Repository) *apiKeyMiddleware {
	return &apiKeyMiddleware{
		repo: repo,
	}
}

func (m apiKeyMiddleware) ApiKeyMiddleware(c fiber.Ctx) error {

	apiKey := c.Get("X-dyno-api-key")

	if apiKey == "" {
		return c.Status(403).JSON(fiber.Map{"error": "Invaild API Key"})
	}

	c.Locals("apiKey", apiKey)

	return c.Next()

}
