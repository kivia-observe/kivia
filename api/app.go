package api

import (
	"github.com/gofiber/fiber/v3"
	apikey "github.com/winnerx0/dyno/internal/api_key"
	"github.com/winnerx0/dyno/internal/auth"
	"github.com/winnerx0/dyno/internal/config"
	"github.com/winnerx0/dyno/internal/database"
	"github.com/winnerx0/dyno/internal/middleware"
	"github.com/winnerx0/dyno/internal/project"
	refreshtoken "github.com/winnerx0/dyno/internal/refresh_token"
	"github.com/winnerx0/dyno/internal/user"
)

type Server struct {
	app *fiber.App

	config config.Config
}

func NewServer(cfg config.Config) *Server {

	app := fiber.New(fiber.Config{
		AppName: "Dyno",
	})

	db := database.Connect(cfg.DBUrl)

	userRepository := user.NewRepository(db)

	refreshTokenRepository := refreshtoken.NewRepository(db)

	userService := auth.NewUserService(*userRepository, *refreshTokenRepository, cfg)

	authhandler := auth.NewAuthHandler(*userService)

	apiRepository := apikey.NewRepository(db)

	apiService := apikey.NewApiKeyService(*apiRepository)

	_ = middleware.NewAuthMiddleware(*apiRepository)

	jwtMiddleware := middleware.NewJwtMiddleware(*userRepository, cfg)

	apiKeyHandler := apikey.NewApiKeyHandler(*apiService)

	
	projectRepository := project.NewRepository(db)

	projectService := project.NewProjectService(*projectRepository)

	projectHandler := project.NewProjectHandler(*projectService)

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hey"})
	})

	v1 := app.Group("api/v1", jwtMiddleware.JwtMiddleware)

	authRouter := app.Group("/auth")

	authRouter.Post("/register", authhandler.Register)

	authRouter.Post("/login", authhandler.Login)

	projectRouter := v1.Group("/projects")
	
	projectRouter.Post("/create", projectHandler.CreateProject)

	apiKeyRouter := v1.Group("/api-keys")

	apiKeyRouter.Post("/create", apiKeyHandler.CreateApiKey)

	apiKeyRouter.Get("/all", apiKeyHandler.GetAllApiKeys)

	apiKeyRouter.Patch("/revoke/:id", apiKeyHandler.RevokeApiKey)

	return &Server{app: app, config: cfg}
}

func (s *Server) Start() error {

	return s.app.Listen(":" + s.config.Port)
}
