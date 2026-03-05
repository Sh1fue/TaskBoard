package repository

import (
	"context"
	"trello_parody/cmd/server/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetById(ctx context.Context, id int) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int) error
}

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context,id int) error
	GetById(ctx context.Context, id int) (*domain.Task, error)
	GetByUserId(ctx context.Context, userId int) (*domain.Task, error)
	GetDeadline(ctx context.Context, userID int) ([]*domain.Task, error)
}
