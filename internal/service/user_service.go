// File: internal/service/user_service.go
package service

import (
	"errors"

	"github.com/veilchrome/myanilog-be/internal/domain"
	"github.com/veilchrome/myanilog-be/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(username, email, password string) (*domain.User, error)
	Login(username, password string) (*domain.User, error)
	LoginWithIdentifier(identifier, password string) (*domain.User, error) // ← tambahkan ini
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(username, email, password string) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:       uuid.NewString(),
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(username, password string) (*domain.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *userService) LoginWithIdentifier(identifier, password string) (*domain.User, error) {
	user, err := s.repo.FindByUsernameOrEmail(identifier)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
