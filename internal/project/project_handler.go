package project

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type projecthandler struct {
	service projectservice
}

func NewProjectHandler(service projectservice) *projecthandler {
	return &projecthandler{
		service: service,
	}
}

func (h projecthandler) CreateProject(c fiber.Ctx) error {

	var projectRequest createProjectRequest

	if err := c.Bind().JSON(&projectRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	project := &Project{
		Id: uuid.New().String(),
		Name: projectRequest.Name,
		UserId: c.Value("userId").(string),
	}

	if err := h.service.CreateProject(project); err != nil {
		if err.Error() == "DUPLICATE_PROJECT_NAME" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Project with name already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Project created successfully"})
}