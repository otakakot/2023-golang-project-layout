package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        string
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateTodo(
	title string,
) *Todo {
	now := time.Now()

	return &Todo{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (td Todo) Complete() *Todo {
	return &Todo{
		ID:        td.ID,
		Title:     td.Title,
		Completed: true,
		CreatedAt: td.CreatedAt,
		UpdatedAt: time.Now(),
	}
}

type TodoList []Todo

func (tl TodoList) Todos() []Todo {
	return tl
}
