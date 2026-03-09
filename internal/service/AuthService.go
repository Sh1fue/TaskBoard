package service

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v5/pgxpool"
	"trello_parody/internal/domain"
	"trello_parody/pkg/jwt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type AuthService struct {
	db         *pgxpool.Pool
	jwtManager *jwt.JWTManager
}

func NewAuthService(db *pgxpool.Pool) *AuthService {
	return &AuthService{
		db:         db,
		jwtManager: jwt.NewJWTManager("secret_key", time.Hour*24),
	}
}

func (s *AuthService) CreateUser(ctx context.Context, username, email, password string)(*domain.User,string,error){
	var id int
	err:= s.db.QueryRow(ctx,"SELECT id FROM users WHERE email=$1",email).Scan(&id)
	if err == nil{
		return nil,"",ErrUserExists
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &domain.User{
        Username: username,
        Email:    email,
        Password: string(hashed),
        CreatedAt: time.Now(),
    }
	err = s.db.QueryRow(ctx,
	"INSERT INTO users (username, email, password, created_at) VALUES ($1,$2,$3,$4) RETURNING id",
	user.Username,user.Email,user.Password,user.CreatedAt,).Scan(&user.ID)
	if err != nil {
		return nil,"",err
	}

	token, _:= s.jwtManager.Generate(user.ID,user.Email)
	return user,token,nil
}

func (s *AuthService) LoginUser(ctx context.Context, email,password string)(*domain.User, string,error){
	user:= &domain.User{}


	query:= `SELECT id,username,email,password,created_at
	FROM users
	WHERE email = $1`

	err := s.db.QueryRow(ctx,query,email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil{
		return nil,"",ErrInvalidCredentials
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil{
		return nil,"",ErrInvalidCredentials
	}

	token,err:= s.jwtManager.Generate(user.ID,user.Email)
	if err != nil{
		return nil,"",err
	}
	return user,token,nil
}