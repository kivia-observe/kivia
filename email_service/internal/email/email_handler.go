package email

import "github.com/gofiber/fiber/v3"

type emailhandler struct {

	emailservice emailservice
}

func NewEmailHandler(emailservice emailservice) *emailhandler {
	return &emailhandler{
		emailservice: emailservice,
	}
}

func (h *emailhandler) SendEmail(c fiber.Ctx) error {
	
	var emailRequest Email

	if err := c.Bind().Body(&emailRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Invalid request"})
	}
	
	err := h.emailservice.SendEmail(emailRequest)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email"})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Email processing"})
}