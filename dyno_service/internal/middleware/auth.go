package middleware

import "github.com/gofiber/fiber/v3"

func AuthMiddlware(c fiber.Ctx) error {

	userId := c.Get("X-User-Id")

	userRole := c.Get("X-User-Role")

	if userId == ""  {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	c.Locals("userId", userId)
	c.Locals("userRole", userRole)

	return c.Next()
}