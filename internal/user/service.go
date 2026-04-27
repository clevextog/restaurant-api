package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo UserRepository
}

func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(email, password string) error {
	err := validateRegisterInput(email, password)
	if err != nil {
		return err
	}
	_, err = s.repo.GetUserByEmail(email)
	if err == nil {
		return ErrUserExists
	}
	if !errors.Is(err, ErrUserNotFound) {
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	passwordHash := string(pass)
	err = s.repo.CreateUser(email, passwordHash)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Login(email, password string) error {
	err := validateLoginInput(email, password)
	if err != nil {
		return err
	}
	user, err := s.repo.GetUserByEmail(email)
	if errors.Is(err, ErrUserNotFound) {
		return ErrInvalidCredentials
	}
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return ErrInvalidCredentials
	}
	return nil
}
