package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/otakakot/2023-golang-project-layout/internal/domain/model"
	"github.com/otakakot/2023-golang-project-layout/internal/domain/repository"
)

var _ repository.Todo = (*Todo)(nil)

type Todo struct {
	db *sql.DB
}

func NewTodo(
	db *sql.DB,
) *Todo {
	return &Todo{
		db: db,
	}
}

// Delete implements repository.Todo.
func (gw *Todo) Delete(
	ctx context.Context,
	id string,
) error {
	query := `DELETE FROM todos WHERE id = $1`

	if _, err := gw.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete row: %w", err)
	}

	return nil
}

// Find implements repository.Todo.
func (gw *Todo) Find(
	ctx context.Context,
	id string,
) (*model.Todo, error) {
	query := `SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = $1`

	row := gw.db.QueryRowContext(ctx, query, id)

	var todo model.Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("todo not found")
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &todo, nil
}

// List implements repository.Todo.
func (gw *Todo) List(ctx context.Context) (model.TodoList, error) {
	query := `SELECT id, title, completed, created_at, updated_at FROM todos`

	rows, err := gw.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query rows: %w", err)
	}
	defer rows.Close()

	todos := make(model.TodoList, 0, 100)
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// Save implements repository.Todo.
func (gw *Todo) Save(
	ctx context.Context,
	todo model.Todo,
) error {
	query := `
		INSERT INTO todos (id, title, completed, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE 
		SET title = $2, completed = $3, updated_at = $5
		`

	if _, err := gw.db.ExecContext(ctx, query, todo.ID, todo.Title, todo.Completed, todo.CreatedAt, todo.UpdatedAt); err != nil {
		return fmt.Errorf("failed to insert row: %w", err)
	}

	return nil
}
