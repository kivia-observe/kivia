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

func (h loghandler) GetLogsByProjectId(c fiber.Ctx) error {

	projectId := c.Params("projectId")

	var startPtr, endPtr *string

	startDate := c.Query("startDate")

	if startDate != "" {
		startPtr = &startDate
	}
	
	endDate := c.Query("endDate")

	if endDate != ""{
		endPtr = &endDate
	}
	


	logs, err := h.service.GetLogsByProjectId(projectId, startPtr, endPtr)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(logs)
}