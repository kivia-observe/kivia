package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/winnerx0/dyno/internal/config"
	"github.com/winnerx0/dyno/internal/database"
	"github.com/winnerx0/dyno/internal/user"
)

type Server struct {

	app *fiber.App

	config config.Config

}

func NewServer(cfg config.Config) *Server{

	app := fiber.New(fiber.Config{
		AppName: "Dyno",

	})

	db := database.Connect(cfg.DBUrl)

	userRepository := user.NewRepository(db)

	userService := user.NewUserService(*userRepository)

	userHandler := user.NewUserHandler(*userService)


	app.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hey"})
	})

	v1 := app.Group("api/v1")

	userRouter := v1.Group("/users")

	userRouter.Post("/register", userHandler.CreateUser)

	return &Server{app: app, config: cfg}
}

func (s *Server) Start() error {

	return s.app.Listen(":" + s.config.Port)
}