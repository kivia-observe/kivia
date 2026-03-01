package apikey

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type apiKeyHandler struct {
	service apiKeyService
}

func NewApiKeyHandler(service apiKeyService) *apiKeyHandler {
	return &apiKeyHandler{
		service: service,
	}
}

func (h apiKeyHandler) CreateApiKey(c fiber.Ctx) error {

	var apiKeyRequest createApiKeyRequest

	if err := c.Bind().JSON(&apiKeyRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	apiKey := &ApiKey{
		Id:        uuid.New().String(),
		Name:      apiKeyRequest.Name,
		ProjectId: apiKeyRequest.ProjectId,
		UserId:    c.Value("userId").(string),
	}

	response, err := h.service.CreateApiKey(apiKey)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h apiKeyHandler) GetAllApiKeys(c fiber.Ctx) error {

	userId := c.Value("userId").(string)

	apiKeys, err := h.service.GetAllApiKeys(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(apiKeys)
}

func (h apiKeyHandler) RevokeApiKey(c fiber.Ctx) error {

	apiKeyId := c.Params("id")
	userId := c.Value("userId").(string)

	err := h.service.RevokeApiKey(apiKeyId, userId)

	if err != nil {

		if errors.Is(err, ErrApiKeyNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, ErrUnauthorized) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "API key revoked successfully"})
}
