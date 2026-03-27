package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/winnerx0/dyno/internal/config"
	refreshtoken "github.com/winnerx0/dyno/internal/refresh_token"
	"github.com/winnerx0/dyno/internal/user"
	"github.com/winnerx0/dyno/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type authservice struct {
	userrepo         *user.Repository
	refreshTokenRepo *refreshtoken.Repository
	config           config.Config
}

func NewAuthService(userrepo *user.Repository, refreshTokenRepo *refreshtoken.Repository, config config.Config) *authservice {
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
		Provider: "local",
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

	sha := sha256.Sum256([]byte(refreshTokenString))
	
	storedRefreshToken := s.refreshTokenRepo.FindByToken(hex.EncodeToString(sha[:]))
	
	if storedRefreshToken.Token == "" {
		return nil, errors.New("Invalid refresh token")
	}
	
	if storedRefreshToken.ExpiresAt.Before(time.Now()) {
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

func (s authservice) generateTokens(userId string) (*AuthResponse, error) {
	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24 * 30)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": refreshTokenExpiresAt.Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(s.config.JwtRefreshTokenSecret))
	if err != nil {
		return nil, err
	}

	accessTokenExpiresAt := time.Now().Add(time.Minute * 15)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
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
		UserId:    userId,
	})
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s authservice) GoogleLogin(code string) (*AuthResponse, error) {
	// Exchange auth code for tokens
	data := url.Values{
		"code":          {code},
		"client_id":     {s.config.GoogleClientID},
		"client_secret": {s.config.GoogleClientSecret},
		"redirect_uri":  {s.config.GoogleRedirectURL},
		"grant_type":    {"authorization_code"},
	}

	resp, err := http.PostForm("https://oauth2.googleapis.com/token", data)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("google token exchange failed: %s", string(body))
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	// Get user info from Google
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tokenResp.AccessToken)

	userResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer userResp.Body.Close()

	var googleUser GoogleUserInfo
	if err := json.NewDecoder(userResp.Body).Decode(&googleUser); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	// Find or create user
	existingUser := s.userrepo.FindByEmail(googleUser.Email)

	if existingUser.Id != "" {
		if existingUser.Provider != "google" {
			return nil, utils.ErrEmailExistsWithDifferentProvider
		}
		return s.generateTokens(existingUser.Id)
	}

	newUser := user.User{
		Id:       uuid.NewString(),
		Email:    googleUser.Email,
		Name:     googleUser.Name,
		ProfilePicture: googleUser.Picture,
		Provider: "google",
		JoinedAt: time.Now(),
	}

	if err := s.userrepo.Save(newUser); err != nil {
		return nil, err
	}

	return s.generateTokens(newUser.Id)
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
