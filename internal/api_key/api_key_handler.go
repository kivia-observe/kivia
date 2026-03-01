package apikey

import (
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
