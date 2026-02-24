package user

import "github.com/gofiber/fiber/v3"

type userhandler struct {
	service userservice
}

func NewUserHandler(service userservice) *userhandler {
	return &userhandler{
		service: service,
	}
}

func (h userhandler) CreateUser(c fiber.Ctx) error {

	var userRequest createUserRequest

	if err := c.Bind().JSON(&userRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateUser(c, userRequest); err != nil {

		if err.Error() == "ALREADY_EXISTS" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User with email already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}
