package service

import (
	"context"
	"fmt"
	"trello_parody/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskService struct {
	db *pgxpool.Pool
}

func NewTaskService(db *pgxpool.Pool) *TaskService {
	return &TaskService{
		db: db,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task *domain.Task) error {
	query := `
	INSERT INTO tasks (user_id, name, description, priority, status, due_date)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id, created_at
	`

	err := s.db.QueryRow(
		ctx,
		query,
		task.UserId,
		task.Name,
		task.Description,
		task.Priority,
		task.Status,
		task.DueDate,
	).Scan(&task.ID, &task.CreatedAt)
	if err != nil {
		return err
	}
	return err
}

func (s *TaskService) UpdateTask(ctx context.Context, task *domain.Task) error {
	query := `
	UPDATE TASKS
	SET name=$1, description=$2, priority=$3,status=$4,due_date=$5,
	WHERE id=$6
	`

	_, err := s.db.Exec(
		ctx,
		query,
		task.Name,
		task.Description,
		task.Priority,
		task.Status,
		task.DueDate,
	)
	if err != nil {
		return fmt.Errorf("Неверные данные")
	}
	return err
}
