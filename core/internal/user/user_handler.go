package user

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

type userhandler struct {
	service userservice
}

func NewUserHandler(service userservice) *userhandler {
	return &userhandler{service: service}
}

func (h *userhandler) GetCurrentUser(c fiber.Ctx) error {
	
	userId := c.Locals("userId").(string)
	
	user, err := h.service.GetUser(userId)
	
	if err != nil {
		if errors.Is(err, UserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *userhandler) DeleteUser(c fiber.Ctx) error {
	
	userId := c.Locals("userId").(string)
	
	err := h.service.DeleteUser(userId)

	if err != nil {
		if errors.Is(err, UserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}

func (h *userhandler) UpdateUser(c fiber.Ctx) error {
	
	userId := c.Locals("userId").(string)
	
	var user editUserRequest
	if err := c.Bind().JSON(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	
	err := h.service.UpdateUser(userId, user)

	if err != nil {
		if errors.Is(err, UserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}