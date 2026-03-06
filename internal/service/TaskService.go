package service

import (
	"errors"
	"time"
	"trello_parody/internal/domain"
)

type CreateTask struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Priority    domain.Priority `json:"priority"`
	Status      domain.Status   `json:"status"`
	DueDate     time.Time       `json:"due_date"`
	Timezone    string          `json:"timezone"`
}
type CreateResponse struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Priority    domain.Priority `json:"priority"`
	Status      domain.Status   `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	DueDate     time.Time       `json:"due_date"`
}

func Create(req *CreateTask) (*CreateResponse, error) {
	if req.Name == "" {
		return nil, errors.New("Поле обязательно к вводу")
	}
	if req.DueDate.Before(time.Now()){
		return nil, errors.New("Дата не может быть указана в прошлом")
	}
	loc, err := time.LoadLocation(req.Timezone)
	if err != nil {
		loc = time.UTC
	}

	return &CreateResponse{
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      req.Status,
		CreatedAt:   time.Now().In(loc),
		DueDate:     req.DueDate,
	},nil

}
