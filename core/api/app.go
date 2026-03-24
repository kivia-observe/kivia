package api

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	_ "github.com/gofiber/fiber/v3/middleware/cors"
	apikey "github.com/winnerx0/dyno/internal/api_key"
	"github.com/winnerx0/dyno/internal/auth"
	"github.com/winnerx0/dyno/internal/config"
	"github.com/winnerx0/dyno/internal/database"
	"github.com/winnerx0/dyno/internal/log"
	"github.com/winnerx0/dyno/internal/middleware"
	"github.com/winnerx0/dyno/internal/project"
	"github.com/winnerx0/dyno/internal/rabbitmq"
	refreshtoken "github.com/winnerx0/dyno/internal/refresh_token"
	"github.com/winnerx0/dyno/internal/sse"
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

	logger := slog.Default()

		server := sse.NewSSEServer(context.Background(), logger)

		server.SrvWg.Add(1)
		go server.Run()

	db := database.Connect(cfg.DBUrl)

	if err := rabbitMQClient.SetupQueues(); err != nil {
		panic(err)
	}

	projectRepository := project.NewRepository(db)

	projectService := project.NewProjectService(projectRepository)

	projectHandler := project.NewProjectHandler(*projectService)

	apiKeyRepository := apikey.NewRepository(db)

	apiService := apikey.NewApiKeyService(apiKeyRepository, projectRepository)

	apiKeyMiddlware := middleware.NewApiKeyMiddleware(apiKeyRepository)

	logRepository := log.NewRepository(db)

	logService := log.NewLogService(logRepository, apiKeyRepository, rabbitMQClient, server)

	logHandler := log.NewLogHandler(*logService)

	refreshTokenRepository := refreshtoken.NewRepository(db)

	userRepository := user.NewRepository(db)

    authService := auth.NewAuthService(userRepository, refreshTokenRepository, cfg)

	_ = middleware.NewJwtMiddleware(*userRepository, cfg)

	apiKeyHandler := apikey.NewApiKeyHandler(*apiService)

	userService := user.NewUserService(*userRepository)

	userHandler := user.NewUserHandler(*userService)

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"http://localhost:3000"},
	// 	AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Dyno-Api-Key", "X-User-ID"},
	// }))

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hey"})
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// auth routes
	authRouter := app.Group("/auth")

	authhandler := auth.NewAuthHandler(*authService)

	authRouter.Post("/register", authhandler.Register)

	authRouter.Post("/login", authhandler.Login)

	authRouter.All("/validate", authhandler.ValidateToken)

	authRouter.Post("/refresh", authhandler.Refresh)

	// project routes
	projectRouter := v1.Group("/projects", middleware.AuthMiddlware)

	projectRouter.Post("/create", projectHandler.CreateProject)

	projectRouter.Get("/all", projectHandler.GetAllProjects)

	// api key routes
	apiKeyRouter := v1.Group("/api-keys", middleware.AuthMiddlware)

	apiKeyRouter.Post("/create", apiKeyHandler.CreateApiKey)

	apiKeyRouter.Get("/all/:projectId", apiKeyHandler.GetAllApiKeys)

	apiKeyRouter.Patch("/revoke/:id", apiKeyHandler.RevokeApiKey)

	// log routes
	logRouter := v1.Group("/logs")

	logRouter.Post("/create", apiKeyMiddlware.ApiKeyMiddleware, logHandler.CreateLog)

	logRouter.Get("/all/:projectId", middleware.AuthMiddlware, logHandler.GetLogsByProjectId)

	logRouter.Get("/stream/:projectId", middleware.AuthMiddlware, server.HandleConnection)

	// user routes
	userRouter := v1.Group("/users", middleware.AuthMiddlware)

	userRouter.Get("/me", userHandler.GetCurrentUser)

	userRouter.Delete("/me", userHandler.DeleteUser)

	userRouter.Put("/me", userHandler.UpdateUser)

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
