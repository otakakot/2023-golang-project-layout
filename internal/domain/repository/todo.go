package repository

import (
	"context"

	"github.com/otakakot/2023-golang-project-layout/internal/domain/model"
)

type Todo interface {
	List(ctx context.Context) (model.TodoList, error)
	Find(ctx context.Context, id string) (*model.Todo, error)
	Save(ctx context.Context, todo model.Todo) error
	Delete(ctx context.Context, id string) error
}
