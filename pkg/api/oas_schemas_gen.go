// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"

	"github.com/google/uuid"
)

// Ref: #/components/schemas/CreateTodoRequest
type CreateTodoRequest struct {
	// The title of the TODO item.
	Title string `json:"title"`
}

// GetTitle returns the value of Title.
func (s *CreateTodoRequest) GetTitle() string {
	return s.Title
}

// SetTitle sets the value of Title.
func (s *CreateTodoRequest) SetTitle(val string) {
	s.Title = val
}

// DeleteTodoNoContent is response for DeleteTodo operation.
type DeleteTodoNoContent struct{}

// Ref: #/components/schemas/Todo
type Todo struct {
	// The ID of the TODO item.
	ID uuid.UUID `json:"id"`
	// The title of the TODO item.
	Title string `json:"title"`
	// Whether the TODO item is completed or not.
	Completed bool `json:"completed"`
	// The date and time when the TODO item was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time when the TODO item was updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// GetID returns the value of ID.
func (s *Todo) GetID() uuid.UUID {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *Todo) GetTitle() string {
	return s.Title
}

// GetCompleted returns the value of Completed.
func (s *Todo) GetCompleted() bool {
	return s.Completed
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Todo) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Todo) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Todo) SetID(val uuid.UUID) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *Todo) SetTitle(val string) {
	s.Title = val
}

// SetCompleted sets the value of Completed.
func (s *Todo) SetCompleted(val bool) {
	s.Completed = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Todo) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Todo) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

// Ref: #/components/schemas/UpdateTodoRequest
type UpdateTodoRequest struct {
	// The title of the TODO item.
	Title string `json:"title"`
	// Whether the TODO item is completed or not.
	Completed bool `json:"completed"`
}

// GetTitle returns the value of Title.
func (s *UpdateTodoRequest) GetTitle() string {
	return s.Title
}

// GetCompleted returns the value of Completed.
func (s *UpdateTodoRequest) GetCompleted() bool {
	return s.Completed
}

// SetTitle sets the value of Title.
func (s *UpdateTodoRequest) SetTitle(val string) {
	s.Title = val
}

// SetCompleted sets the value of Completed.
func (s *UpdateTodoRequest) SetCompleted(val bool) {
	s.Completed = val
}