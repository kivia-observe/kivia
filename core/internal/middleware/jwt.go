package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/winnerx0/kivia/internal/config"
	"github.com/winnerx0/kivia/internal/user"
)

type jwtMiddleware struct {
	userRepo user.Repository
	config   config.Config
}

func NewJwtMiddleware(userRepo user.Repository, config config.Config) *jwtMiddleware {
	return &jwtMiddleware{
		userRepo: userRepo,
		config:   config,
	}
}

type CustomCLaims struct {
	role string

	jwt.RegisteredClaims
}

func (m jwtMiddleware) JwtMiddleware(c fiber.Ctx) error {

	authHeader := c.Get(fiber.HeaderAuthorization)

	if authHeader == "" || authHeader[:7] != "Bearer " {
		return c.Status(401).JSON(fiber.Map{"error": "Missing or invalid Authorization header"})
	}

	authToken := authHeader[7:]

	var claims = &CustomCLaims{}

	token, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (any, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}

		return []byte(m.config.JwtAccessTokenSecret), nil
	})

	if err != nil || !token.Valid {

		return c.Status(401).JSON(fiber.Map{"error": "Invalid Token"})
	}

	user := m.userRepo.FindById(claims.Subject)

	if user.Id == "" {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	c.Locals("userId", user.Id)

	return c.Next()

}
