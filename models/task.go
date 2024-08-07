package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Status represents the possible statuses of a task.
type Status string

const (
	StatusDone       Status = "done"
	StatusInProgress Status = "inprogress"
	StatusPending    Status = "pending"
)

// Task represents a task with title, description, due date, status, and ID.
type Task struct {
	id          uuid.UUID
	title       string
	description string
	dueDate     time.Time
	status      Status
}

// TaskBSON represents the BSON format of a Task for MongoDB operations.
type TaskBSON struct {
	ID          uuid.UUID `bson:"_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	DueDate     time.Time `bson:"dueDate"`
	Status      Status    `bson:"status"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}

// ToBSON converts a Task to a TaskBSON.
func (t *Task) ToBSON() *TaskBSON {
	return &TaskBSON{
		ID:          t.ID(),
		Title:       t.Title(),
		Description: t.Description(),
		DueDate:     t.DueDate(),
		Status:      t.Status(),
		UpdatedAt:   time.Now(), // Update timestamp on modification
	}
}

// FromBSON converts a TaskBSON to a Task.
func FromBSON(bson *TaskBSON) *Task {
	return &Task{
		id:          bson.ID,
		title:       bson.Title,
		description: bson.Description,
		dueDate:     bson.DueDate,
		status:      bson.Status,
	}
}

// TaskConfig represents the configuration for creating or updating a Task.
type TaskConfig struct {
	Title       string
	Description string
	DueDate     time.Time
	Status      Status
}

// NewTask creates a new Task with the given configuration, validates its properties, and generates an ID.
func NewTask(config TaskConfig) (*Task, error) {
	if err := validateTaskConfig(config); err != nil {
		return nil, err
	}

	return &Task{
		id:          uuid.New(),
		title:       config.Title,
		description: config.Description,
		dueDate:     config.DueDate,
		status:      config.Status,
	}, nil
}

// validateTaskConfig checks if the provided task configuration is valid.
func validateTaskConfig(config TaskConfig) error {
	if config.Title == "" {
		return errors.New("title cannot be empty")
	}
	if config.Description == "" {
		return errors.New("description cannot be empty")
	}
	if config.DueDate.IsZero() {
		return errors.New("due date cannot be zero")
	}
	if !isValidStatus(config.Status) {
		return errors.New("invalid status")
	}
	return nil
}

// isValidStatus checks if the given status is one of the allowed statuses.
func isValidStatus(status Status) bool {
	switch status {
	case StatusDone, StatusInProgress, StatusPending:
		return true
	default:
		return false
	}
}

// ID returns the task's ID.
func (t *Task) ID() uuid.UUID {
	return t.id
}

// Title returns the task's title.
func (t *Task) Title() string {
	return t.title
}

// Description returns the task's description.
func (t *Task) Description() string {
	return t.description
}

// DueDate returns the task's due date.
func (t *Task) DueDate() time.Time {
	return t.dueDate
}

// Status returns the task's status.
func (t *Task) Status() Status {
	return t.status
}

// Update updates the task's fields with the provided configuration after validating the data.
func (t *Task) Update(config TaskConfig) error {
	if err := validateTaskConfig(config); err != nil {
		return err
	}

	t.title = config.Title
	t.description = config.Description
	t.dueDate = config.DueDate
	t.status = config.Status
	return nil
}

