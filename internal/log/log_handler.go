package log

import "github.com/gofiber/fiber/v3"

type loghandler struct {
	service Logservice
}

func NewLogHandler(service Logservice) *loghandler {
	return &loghandler{
		service: service,
	}
}

func (h loghandler) CreateLog(c fiber.Ctx) error {
	
	projectId := c.Params("projectId")

	var createLogRequest createLogRequest

	if err := c.Bind().JSON(&createLogRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateLog(createLogRequest, projectId); err != nil {
	
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Log sent to queue"})
}