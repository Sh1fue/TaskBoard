package service

import (
	"context"
	"errors"
	"time"
	"trello_parody/internal/domain"
	"trello_parody/internal/repository"
	"trello_parody/pkg/jwt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type AuthService struct {
	UserRepo   repository.UserRepository
	jwtManager *jwt.JWTManager
}

func NewAuthService(userRepo repository.UserRepository, jwtManager *jwt.JWTManager) *AuthService {
	return &AuthService{
		UserRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthResponse struct {
	Token string       `json:"token"`
	User  *domain.User `json:"user"`
}

func (s *AuthService) Register(ctx context.Context, req *RegisterRequest) (*AuthResponse, error) {
	if _, err := s.UserRepo.GetByEmail(ctx, req.Email); err == nil {
		return nil, ErrUserExists
	}

	user := &domain.User{
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}
	if err := user.SetPassword(req.Password); err != nil {
		return nil, err
	}

	if err := s.UserRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	token, err := s.jwtManager.Generate(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error) {
	{

		user, err := s.UserRepo.GetByEmail(ctx, req.Email)
		if err != nil {
			return nil, ErrInvalidCredentials
		}

		if !user.CheckPassword(req.Password) {
			return nil, ErrInvalidCredentials
		}

		token, err := s.jwtManager.Generate(user.ID, user.Email)
		if err != nil {
			return nil, err
		}

		return &AuthResponse{
			Token: token,
			User:  user,
		}, nil
	}
}
