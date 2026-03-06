package service

import (
    "context"
    "errors"
    "time"

    "trello_parody/internal/domain"

    "github.com/jackc/pgx/v5/pgxpool"
    "trello_parody/pkg/jwt"
    "golang.org/x/crypto/bcrypt"
)

var (
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrUserExists         = errors.New("user already exists")
)

type AuthService struct {
    db         *pgxpool.Pool
    jwtManager *jwt.JWTManager
}

func NewAuthService(db *pgxpool.Pool, jwtManager *jwt.JWTManager) *AuthService {
    return &AuthService{
        db:         db,
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
    var existingID int
    err := s.db.QueryRow(ctx, "SELECT id FROM users WHERE email=$1", req.Email).Scan(&existingID)
    if err == nil {
        return nil, ErrUserExists
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &domain.User{
        Username:  req.Username,
        Email:     req.Email,
        Password:  string(hashed),
        CreatedAt: time.Now(),
    }

    // Вставка в БД и получение ID и CreatedAt
    query := `
        INSERT INTO users (username, email, password, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
    err = s.db.QueryRow(ctx, query, user.Username, user.Email, user.Password, user.CreatedAt).Scan(&user.ID)
    if err != nil {
        return nil, err
    }

    // Генерация JWT
    token, err := s.jwtManager.Generate(user.ID, user.Email)
    if err != nil {
        return nil, err
    }

    return &AuthResponse{
        Token: token,
        User:  user,
    }, nil
}

// --------------------- LOGIN ---------------------
func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error) {
    user := &domain.User{}
    query := "SELECT id, username, email, password, created_at FROM users WHERE email=$1"
    err := s.db.QueryRow(ctx, query, req.Email).Scan(
        &user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
    )
    if err != nil {
        return nil, ErrInvalidCredentials
    }

    // Проверка пароля
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
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