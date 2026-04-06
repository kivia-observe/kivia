package api

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/winnerx0/kivia/email_service/internal/config"
	"github.com/winnerx0/kivia/email_service/internal/email"
	"github.com/winnerx0/kivia/email_service/internal/rabbitmq"
)

type Server struct {
	app *fiber.App

	config config.Config
	cancel context.CancelFunc
}

func NewServer(config config.Config, rabbitMQClient *rabbitmq.RabbitMQClient) *Server {

	app := fiber.New(fiber.Config{
		AppName: "Dyno - Email Service",
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	emailService := email.NewEmailService(config.BrevoApiKey, config.SenderEmail, config.SenderName, rabbitMQClient)

	emailHandler := email.NewEmailHandler(*emailService)

	api := app.Group("/api")

	v1 := api.Group("/v1")

	emailRouter := v1.Group("/email")

	emailRouter.Post("/send", emailHandler.SendEmail)

	ctx, cancel := context.WithCancel(context.Background())

	if err := emailService.EmailConsumer(ctx); err != nil {
		slog.Error("Error starting log consumer: ", err)
	}

	return &Server{app: app, config: config, cancel: cancel}
}

func (s *Server) Start() error {

	defer s.cancel()

	return s.app.Listen(":" + s.config.Port)
}
