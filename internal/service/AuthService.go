package service

import "trello_parody/cmd/server/internal/repository"

type AuthService struct {
	UserRepo   repository.UserRepository
	jwtManager *jwt.JWTManager
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json"password"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json"password"`
}
