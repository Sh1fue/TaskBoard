package domain

import (
	"time"
)

type Priority int

const (
	PrioritySoSo Priority = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
	PriorityCritical
)

func (p Priority) String() string {
	switch p {
	case PrioritySoSo:
		return "Можно не торопиться"
	case PriorityLow:
		return "Низкий"
	case PriorityMedium:
		return "Средний"
	case PriorityHigh:
		return "Высокий"
	case PriorityCritical:
		return "Критический"
	default:
		return "Неизвестно"
	}
}

type Status string

const (
	StatusNotActive Status = "Не взяты"
	StatusActive    Status = "Активны"
	StatusDone      Status = "Готовы"
	StatusDeadLine  Status = "Сгорели"
)

type Task struct {
	ID          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	DueDate     time.Time `json:"due_date"`
}

func (t *Task) IsDead() bool {
	if t.Status == StatusDone {
		return false
	}
	return !t.DueDate.IsZero() && t.DueDate.Before(time.Now())
}
func (t *Task) UpdateStatus() {
	if t.Status == StatusDone {
		return
	}
	if t.IsDead() {
		t.Status = StatusDeadLine
	} else {
		t.Status = StatusActive
	}

}
func (t *Task) IsDueSoon(HoursBefore int) bool {
	if t.Status == StatusDone || t.DueDate.IsZero() {
		return false
	}
	now := time.Now()

	return t.DueDate.Before(now) &&
		t.DueDate.Sub(now).Hours() <= float64(HoursBefore)
}
