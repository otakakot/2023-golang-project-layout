package controller

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/otakakot/2023-golang-project-layout/internal/application/usecase"
	"github.com/otakakot/2023-golang-project-layout/pkg/api"
)

type Todo struct {
	uc usecase.Todo
}

func NewTodo(
	uc usecase.Todo,
) *Todo {
	return &Todo{
		uc: uc,
	}
}

// CreateTodo implements api.Handler.
func (ctl *Todo) CreateTodo(ctx context.Context, req *api.CreateTodoRequest) (*api.Todo, error) {
	output, err := ctl.uc.Create(ctx, usecase.TodoCreateInput{
		Title: req.Title,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	id, _ := uuid.Parse(output.Todo.ID)

	slog.Info(fmt.Sprintf("created todo: %+v", output.Todo))

	return &api.Todo{
		ID:        id,
		Title:     output.Todo.Title,
		Completed: output.Todo.Completed,
		CreatedAt: output.Todo.CreatedAt,
		UpdatedAt: output.Todo.UpdatedAt,
	}, nil
}

// DeleteTodo implements api.Handler.
func (ctl *Todo) DeleteTodo(ctx context.Context, params api.DeleteTodoParams) error {
	if _, err := ctl.uc.Delete(ctx, usecase.TodoDeleteInput{
		ID: params.ID,
	}); err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	return nil
}

// GetTodo implements api.Handler.
func (ctl *Todo) GetTodo(ctx context.Context, params api.GetTodoParams) (*api.Todo, error) {
	output, err := ctl.uc.Read(ctx, usecase.TodoReadInput{
		ID: params.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read todo: %w", err)
	}

	id, _ := uuid.Parse(output.Todo.ID)

	return &api.Todo{
		ID:        id,
		Title:     output.Todo.Title,
		Completed: output.Todo.Completed,
		CreatedAt: output.Todo.CreatedAt,
		UpdatedAt: output.Todo.UpdatedAt,
	}, nil
}

// ListTodos implements api.Handler.
func (ctl *Todo) ListTodos(ctx context.Context) ([]api.Todo, error) {
	output, err := ctl.uc.List(ctx, usecase.TodoListInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list todos: %w", err)
	}

	todos := make([]api.Todo, len(output.TodoList))
	for i, todo := range output.TodoList {
		id, _ := uuid.Parse(todo.ID)

		todos[i] = api.Todo{
			ID:        id,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		}
	}

	return todos, nil
}

// UpdateTodo implements api.Handler.
func (ctl *Todo) UpdateTodo(ctx context.Context, req *api.UpdateTodoRequest, params api.UpdateTodoParams) (*api.Todo, error) {
	output, err := ctl.uc.Update(ctx, usecase.TodoUpdateInput{
		ID:        params.ID,
		Title:     req.Title,
		Completed: req.Completed,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update todo: %w", err)
	}

	id, _ := uuid.Parse(output.Todo.ID)

	return &api.Todo{
		ID:        id,
		Title:     output.Todo.Title,
		Completed: output.Todo.Completed,
		CreatedAt: output.Todo.CreatedAt,
		UpdatedAt: output.Todo.UpdatedAt,
	}, nil
}
