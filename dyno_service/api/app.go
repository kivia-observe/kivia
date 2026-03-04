package api

import (
	"context"

	"github.com/gofiber/fiber/v3"
	apikey "github.com/winnerx0/dyno/internal/api_key"
	"github.com/winnerx0/dyno/internal/auth"
	"github.com/winnerx0/dyno/internal/config"
	"github.com/winnerx0/dyno/internal/database"
	"github.com/winnerx0/dyno/internal/log"
	"github.com/winnerx0/dyno/internal/middleware"
	"github.com/winnerx0/dyno/internal/project"
	"github.com/winnerx0/dyno/internal/rabbitmq"
	refreshtoken "github.com/winnerx0/dyno/internal/refresh_token"
	"github.com/winnerx0/dyno/internal/user"
)

type Server struct {
	app *fiber.App

	config config.Config

	logService log.Logservice
}

func NewServer(cfg config.Config, rabbitMQClient *rabbitmq.RabbitMQClient) *Server {

	app := fiber.New(fiber.Config{
		AppName: "Dyno",
	})

	db := database.Connect(cfg.DBUrl)

	if err := rabbitMQClient.SetupQueues(); err != nil {
		panic(err)
	}

	logRepository := log.NewRepository(db)

	logService := log.NewLogService(*logRepository, *rabbitMQClient)

	logHandler := log.NewLogHandler(*logService)

	userRepository := user.NewRepository(db)

	refreshTokenRepository := refreshtoken.NewRepository(db)

	userService := auth.NewUserService(*userRepository, *refreshTokenRepository, cfg)

	projectRepository := project.NewRepository(db)

	projectService := project.NewProjectService(*projectRepository)

	projectHandler := project.NewProjectHandler(*projectService)

	authhandler := auth.NewAuthHandler(*userService)

	apiRepository := apikey.NewRepository(db)

	apiService := apikey.NewApiKeyService(*apiRepository, *projectRepository)

	apiKeyMiddlware := middleware.NewApiKeyMiddleware(*apiRepository)

	_ = middleware.NewJwtMiddleware(*userRepository, cfg)

	apiKeyHandler := apikey.NewApiKeyHandler(*apiService)

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hey"})
	})

	v1 := app.Group("api/v1", middleware.AuthMiddlware)

	// auth routes
	authRouter := app.Group("/auth")

	authRouter.Post("/register", authhandler.Register)

	authRouter.Post("/login", authhandler.Login)

	authRouter.All("/validate", authhandler.ValidateToken)

	authRouter.Post("/refresh", authhandler.Refresh)

	// project routes
	projectRouter := v1.Group("/projects")

	projectRouter.Post("/create", projectHandler.CreateProject)

	projectRouter.Get("/all", projectHandler.GetAllProjects)

	// api key routes
	apiKeyRouter := v1.Group("/api-keys")

	apiKeyRouter.Post("/create", apiKeyHandler.CreateApiKey)

	apiKeyRouter.Get("/all/:projectId", apiKeyHandler.GetAllApiKeys)

	apiKeyRouter.Patch("/revoke/:id", apiKeyHandler.RevokeApiKey)

	// log routes
	logRouter := v1.Group("/logs/:projectId", apiKeyMiddlware.ApiKeyMiddleware)

	logRouter.Post("/create", logHandler.CreateLog)

	logRouter.Get("/all", logHandler.GetLogsByProjectId)

	return &Server{app: app, config: cfg, logService: *logService}
}

func (s *Server) Start() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.logService.LogConsumer(ctx); err != nil {
		return err
	}

	return s.app.Listen(":" + s.config.Port)
}
