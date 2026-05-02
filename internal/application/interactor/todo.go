package interactor

import (
	"context"
	"fmt"

	"github.com/otakakot/2023-golang-project-layout/internal/application/usecase"
	"github.com/otakakot/2023-golang-project-layout/internal/domain/model"
	"github.com/otakakot/2023-golang-project-layout/internal/domain/repository"
)

var _ usecase.Todo = (*Todo)(nil)

type Todo struct {
	todoRepository repository.Todo
}

func NewTodo(
	todoRepository repository.Todo,
) *Todo {
	return &Todo{
		todoRepository: todoRepository,
	}
}

// Create implements usecase.Todo.
func (itr *Todo) Create(
	ctx context.Context,
	input usecase.TodoCreateInput,
) (*usecase.TodoCreateOutput, error) {
	todo := model.GenerateTodo(input.Title)

	if err := itr.todoRepository.Save(ctx, *todo); err != nil {
		return nil, fmt.Errorf("failed to save todo: %w", err)
	}

	return &usecase.TodoCreateOutput{
		Todo: *todo,
	}, nil
}

// Delete implements usecase.Todo.
func (itr *Todo) Delete(
	ctx context.Context,
	input usecase.TodoDeleteInput,
) (*usecase.TodoDeleteOutput, error) {
	if err := itr.todoRepository.Delete(ctx, input.ID); err != nil {
		return nil, fmt.Errorf("failed to delete todo: %w", err)
	}

	return &usecase.TodoDeleteOutput{}, nil
}

// List implements usecase.Todo.
func (itr *Todo) List(
	ctx context.Context,
	input usecase.TodoListInput,
) (*usecase.TodoListOutput, error) {
	todos, err := itr.todoRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list todos: %w", err)
	}

	return &usecase.TodoListOutput{
		TodoList: todos,
	}, nil
}

// Read implements usecase.Todo.
func (itr *Todo) Read(
	ctx context.Context,
	input usecase.TodoReadInput,
) (*usecase.TodoReadOutput, error) {
	todo, err := itr.todoRepository.Find(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find todo: %w", err)
	}

	return &usecase.TodoReadOutput{
		Todo: *todo,
	}, nil
}

// Update implements usecase.Todo.
func (itr *Todo) Update(
	ctx context.Context,
	input usecase.TodoUpdateInput,
) (*usecase.TodoUpdateOutput, error) {
	todo, err := itr.todoRepository.Find(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find todo: %w", err)
	}

	updated := todo.Complete()

	if err := itr.todoRepository.Save(ctx, *updated); err != nil {
		return nil, fmt.Errorf("failed to save todo: %w", err)
	}

	return &usecase.TodoUpdateOutput{
		Todo: *todo,
	}, nil
}
