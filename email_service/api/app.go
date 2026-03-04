package api

import (

	"github.com/gofiber/fiber/v3"
	"github.com/winnerx0/dyno/email_service/internal/email"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {

	app := fiber.New(fiber.Config{
		AppName: "Dyno - Email Service",
	})

	

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hey"})
	})

	emailService := email.NewEmailService()

	emailHandler := email.NewEmailHandler(*emailService)

	// v1 := app.Group("/email")

	emailRouter := app.Group("/email")

	emailRouter.Post("send",  emailHandler.SendEmail)

	return &Server{app: app}
}

func (s *Server) Start() error {

	return s.app.Listen(":" + "8082")
}
