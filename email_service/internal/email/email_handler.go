package email

import "github.com/gofiber/fiber/v3"

type emailhandler struct {

	emailservice emailservice
}

func NewEmailHandler(emailservice emailservice) *emailhandler {
	return &emailhandler{}
}

func (h *emailhandler) SendEmail(c fiber.Ctx) error {
	type request struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	var req request

	if err := c.JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	to := req.To
	subject := req.Subject
	body := req.Body
	return h.emailservice.SendEmail(to, subject, body)
}