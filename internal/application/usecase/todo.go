package usecase

import (
	"context"

	"github.com/otakakot/2023-golang-project-layout/internal/domain/model"
)

type Todo interface {
	Create(ctx context.Context, input TodoCreateInput) (*TodoCreateOutput, error)
	Read(ctx context.Context, input TodoReadInput) (*TodoReadOutput, error)
	List(ctx context.Context, input TodoListInput) (*TodoListOutput, error)
	Update(ctx context.Context, input TodoUpdateInput) (*TodoUpdateOutput, error)
	Delete(ctx context.Context, input TodoDeleteInput) (*TodoDeleteOutput, error)
}

type TodoCreateInput struct {
	Title string
}

type TodoCreateOutput struct {
	Todo model.Todo
}

type TodoReadInput struct {
	ID string
}

type TodoReadOutput struct {
	Todo model.Todo
}

type TodoListInput struct {
}

type TodoListOutput struct {
	TodoList model.TodoList
}

type TodoUpdateInput struct {
	ID        string
	Title     string
	Completed bool
}

type TodoUpdateOutput struct {
	Todo model.Todo
}

type TodoDeleteInput struct {
	ID string
}

type TodoDeleteOutput struct {
}
