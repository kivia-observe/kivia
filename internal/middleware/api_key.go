package middleware

import (
	"log"

	"github.com/gofiber/fiber/v3"
	apikey "github.com/winnerx0/dyno/internal/api_key"
	"golang.org/x/crypto/bcrypt"
)

type apiKeyMiddleware struct {
	repo apikey.Repository
}

func NewApiKeyMiddleware(repo apikey.Repository) *apiKeyMiddleware {
	return &apiKeyMiddleware{
		repo: repo,
	}
}

func (m apiKeyMiddleware) ApiKeyMiddleware(c fiber.Ctx) error {

	apiKey := c.Get("x-dyno-api-key")

	projectId := c.Params("projectID")

	if apiKey == "" {
		return c.Status(403).JSON(fiber.Map{"error": "Invaild API Key"})
	}

	key, err := m.repo.FindByProjectId(projectId)

	if err != nil {
		log.Println("Error ", err)
		return c.Status(403).JSON(fiber.Map{"error": "Invaild API Key"})
	}

	if key.Id == "" {
		return c.Status(403).JSON(fiber.Map{"error": "Invaild API Key"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(key.Key), []byte(apiKey))

	if err != nil {
		return c.Status(403).JSON(fiber.Map{"error": "Invaild API Key"})
	}

	c.Locals("apiKey", key)

	return c.Next()

}
