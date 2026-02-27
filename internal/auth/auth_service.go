package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/winnerx0/dyno/internal/config"
	refreshtoken "github.com/winnerx0/dyno/internal/refresh_token"
	"github.com/winnerx0/dyno/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type authservice struct {
	userrepo         user.Repository
	refreshTokenRepo refreshtoken.Repository
	config           config.Config
}

func NewUserService(userrepo user.Repository, refreshTokenRepo refreshtoken.Repository, config config.Config) *authservice {
	return &authservice{
		userrepo:         userrepo,
		refreshTokenRepo: refreshTokenRepo,
		config:           config,
	}
}

func (s authservice) Register(createUserRequest createUserRequest) error {

	if exists := s.userrepo.ExistsByEmail(createUserRequest.Email); exists {
		return errors.New("ALREADY_EXISTS")
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), 10)

	if err != nil {
		return err
	}

	user := user.User{
		Id:       uuid.NewString(),
		Email:    createUserRequest.Email,
		Password: string(passwordBytes),
		Name:     createUserRequest.Name,
		JoinedAt: time.Now(),
	}

	if err := s.userrepo.Save(user); err != nil {
		return err
	}

	return nil
}

func (s authservice) Login(loginRequest LoginRequest) (*AuthResponse, error) {

	user := s.userrepo.FindByEmail(loginRequest.Email)

	if user.Id == "" {
		return nil, errors.New("NOT_EXISTS")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {
		return nil, err
	}

	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24 * 30)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": refreshTokenExpiresAt.Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(s.config.JwtRefreshTokenSecret))

	if err != nil {
		return nil, err
	}

	accessTokenExpiresAt := time.Now().Add(time.Minute * 15)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.Id,
		"exp":   accessTokenExpiresAt.Unix(),
		"roles": []string{"USER"},
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.config.JwtAccessTokenSecret))

	if err != nil {
		return nil, err
	}

	err = s.refreshTokenRepo.Save(refreshtoken.RefreshToken{
		Id:        uuid.NewString(),
		Token:     refreshTokenString,
		ExpiresAt: refreshTokenExpiresAt,
		UserId:    user.Id,
	})

	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
