package auth

import "github.com/gofiber/fiber/v3"

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

	if err := h.service.Register(c, userRequest); err != nil {

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

	 response, err := h.service.Login(c, loginRequest);

	if err != nil {

		if err.Error() == "NOT_EXISTS" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
