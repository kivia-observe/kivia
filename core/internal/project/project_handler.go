package project

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/winnerx0/kivia/internal/validator"
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
	
	if err := validator.Get().Struct(projectRequest); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": validator.FirstError(err)})
	}

	project := &Project{
		Id: uuid.New().String(),
		Name: projectRequest.Name,
		UserId: c.Value("userId").(string),
		ApiKeys:[]string{},
	}

	if err := h.service.CreateProject(project); err != nil {
		if err.Error() == "DUPLICATE_PROJECT_NAME" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Project with name already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Project created successfully"})
}

func (h projecthandler) GetAllProjects(c fiber.Ctx) error {

	userId := c.Value("userId").(string)

	projects, err := h.service.GetAllProjectsByUser(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(projects)
}