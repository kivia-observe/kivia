package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
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

type CustomClaims struct {
	role string

	jwt.RegisteredClaims
}

type ValidateResponse struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
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

	hash := sha256.Sum256([]byte(refreshTokenString))

	err = s.refreshTokenRepo.Save(refreshtoken.RefreshToken{
		Id:        uuid.NewString(),
		Token:     hex.EncodeToString(hash[:]),
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

func (s authservice) RefreshToken(refreshTokenString string) (*AuthResponse, error) {

	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(s.config.JwtRefreshTokenSecret), nil
	})

	if err != nil || !refreshToken.Valid {
		return nil, errors.New("Invalid refresh token")
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)

	if !ok || claims["sub"] == nil {
		return nil, errors.New("Invalid refresh token claims")
	}

	userId := claims["sub"].(string)

	storedRefreshToken := s.refreshTokenRepo.FindByToken(refreshTokenString)

	storedHash := sha256.Sum256([]byte(storedRefreshToken.Token))

	sha := sha256.Sum256([]byte(refreshTokenString))

	if !bytes.Equal(storedHash[:], sha[:]) || storedRefreshToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("Invalid or expired refresh token")
	}

	accessTokenExpiresAt := time.Now().Add(time.Minute * 15)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"exp":   accessTokenExpiresAt.Unix(),
		"roles": []string{"USER"},
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.config.JwtAccessTokenSecret))

	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24 * 30)

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": refreshTokenExpiresAt.Unix(),
	})

	newRefreshTokenString, err := newRefreshToken.SignedString([]byte(s.config.JwtRefreshTokenSecret))

	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256([]byte(newRefreshTokenString))

	err = s.refreshTokenRepo.Save(refreshtoken.RefreshToken{
		Id:        uuid.NewString(),
		Token:     hex.EncodeToString(hash[:]),
		ExpiresAt: refreshTokenExpiresAt,
		UserId:    userId,
	})

	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken:  accessTokenString,
		RefreshToken: newRefreshTokenString,
	}, nil
}

func (s authservice) Validate(authToken string) (ValidateResponse, error) {

	var claims = &CustomClaims{}

	token, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (any, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return ValidateResponse{}, errors.New("Unexpected signing method")
		}

		return []byte(s.config.JwtAccessTokenSecret), nil
	})

	if err != nil || !token.Valid {

		return ValidateResponse{}, errors.New("Invalid Token")
	}

	user := s.userrepo.FindById(claims.Subject)

	if user.Id == "" {
		return ValidateResponse{}, errors.New("User not found")
	}

	return ValidateResponse{
		UserId: user.Id,
		Role:   user.Role,
	}, nil
}
