package service

import (
	"errors"
	"fmt"
	"pbkk-fp-go/dto"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/helper"
	"pbkk-fp-go/repository"
)

// AuthService defines the authentication service interface.
type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

// Register registers a new user.
func (s *authService) Register(req *dto.RegisterRequest) error {
	// Validate input
	if req.Name == "" || req.Email == "" || req.Password == "" || req.PasswordConfirmation == "" {
		return fmt.Errorf("all fields are required")
	}

	// Check if email already exists
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return fmt.Errorf("email already registered")
	}

	// Check if passwords match
	if req.Password != req.PasswordConfirmation {
		return fmt.Errorf("passwords do not match")
	}

	// Hash password
	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user entity
	user := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
	}

	// Save user to repository
	if err := s.repository.Register(&user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

// Login authenticates a user and returns a token.
func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Validate input
	if req.Email == "" || req.Password == "" {
		return nil, fmt.Errorf("email and password are required")
	}

	// Find user by email
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("wrong email or password") // Do not expose exact issue
	}

	// Verify password
	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, errors.New("wrong email or password")
	}

	// Generate token
	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Build response
	data := dto.LoginResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Token: token,
	}

	return &data, nil
}
