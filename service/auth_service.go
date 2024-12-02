package service

import (
	"pbkk-fp-go/dto"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/errorhandler"
	"pbkk-fp-go/helper"
	"pbkk-fp-go/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "email already registered"}
	}

	if req.Password != req.PasswordConfirmation {
		return &errorhandler.BadRequestError{Message: "password not match"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Token: token,
	}

	return &data, nil
}