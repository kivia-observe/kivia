package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userservice struct {
	repo repository
}

func NewUserService(repo repository) *userservice {
	return &userservice{
		repo: repo,
	}
}

func (s userservice) CreateUser(ctx context.Context, createUserRequest createUserRequest) error {

	if exists := s.repo.ExistsByEmail(createUserRequest.Email); exists {
		return errors.New("ALREADY_EXISTS")
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), 10)

	if err != nil {
		return err
	}

	user := User{
		Id:       uuid.NewString(),
		Email:    createUserRequest.Email,
		Password: string(passwordBytes),
		Name:     createUserRequest.Name,
		JoinedAt: time.Now(),
	}

	if err := s.repo.Save(user); err != nil {
		return err
	}

	return nil
}
