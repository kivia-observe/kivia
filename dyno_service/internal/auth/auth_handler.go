package auth

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type authhandler struct {
	service authservice
}

func NewAuthHandler(service authservice) *authhandler {
	return &authhandler{
		service: service,
	}
}

func (h authhandler) Register(c fiber.Ctx) error {

	var userRequest createUserRequest

	if err := c.Bind().JSON(&userRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.Register(userRequest); err != nil {

		if err.Error() == "ALREADY_EXISTS" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User with email already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (h authhandler) Login(c fiber.Ctx) error {

	var loginRequest LoginRequest

	if err := c.Bind().JSON(&loginRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	 response, err := h.service.Login(loginRequest);

	if err != nil {

		if err.Error() == "NOT_EXISTS" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h authhandler) Refresh(c fiber.Ctx) error {

	var refreshRequest RefreshRequest

	if err := c.Bind().JSON(&refreshRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := h.service.RefreshToken(refreshRequest.RefreshToken)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h authhandler) ValidateToken(c fiber.Ctx) error {

	authHeader := c.Get(fiber.HeaderAuthorization)

	
	if authHeader == "" || authHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid Authorization header"})
	}

	authToken := authHeader[7:]

	response, err := h.service.Validate(authToken)

	log.Printf("ValidateToken response: %+v, error: %v", response, err)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("X-User-Id", response.UserId)
	c.Set("X-User-Role", response.Role)

	return c.Status(fiber.StatusOK).JSON("OK")
}