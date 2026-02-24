package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/winnerx0/dyno/internal/config"
)

type Server struct {

	app *fiber.App

	config config.Config

}

func NewServer(cfg config.Config) *Server{

	app := fiber.New(fiber.Config{
		AppName: "Dyno",

	})

	return &Server{app: app, config: cfg}
}

func (s *Server) Start() error {

	return s.app.Listen(":" + s.config.Port)
}