package mocks

import (
	"context"

	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

type MockDatabase struct {
	InsertTasksFunc         func(result tasks_domain.Task, collection string) error
	SelectTasksFunc         func(collection string) ([]tasks_domain.Task, error)
	SelectTasksStatusFunc   func(collection, status string) ([]tasks_domain.Task, error)
	SelectTasksDateFunc     func(collection, date string) ([]tasks_domain.Task, error)
	SelectTasksPriorityFunc func(collection, priority string) ([]tasks_domain.Task, error)
	UpdateTasksFunc         func(result tasks_domain.Task, collection string) error
	DeleteTasksFunc         func(result tasks_domain.Task, collection string) error
	TaskMigrationFunc       func(ctx context.Context) error
}

func (m *MockDatabase) InsertTasks(result tasks_domain.Task, collection string) error {
	if m.InsertTasksFunc != nil {
		return m.InsertTasksFunc(result, collection)
	}
	return nil
}

func (m *MockDatabase) SelectTasks(collection string) ([]tasks_domain.Task, error) {
	if m.SelectTasksFunc != nil {
		return m.SelectTasksFunc(collection)
	}
	return nil, nil
}

func (m *MockDatabase) SelectTasksStatus(collection, status string) ([]tasks_domain.Task, error) {
	if m.SelectTasksStatusFunc != nil {
		return m.SelectTasksStatusFunc(collection, status)
	}
	return nil, nil
}

func (m *MockDatabase) SelectTasksDate(collection, date string) ([]tasks_domain.Task, error) {
	if m.SelectTasksDateFunc != nil {
		return m.SelectTasksDateFunc(collection, date)
	}
	return nil, nil
}

func (m *MockDatabase) SelectTasksPriority(collection, priority string) ([]tasks_domain.Task, error) {
	if m.SelectTasksPriorityFunc != nil {
		return m.SelectTasksPriorityFunc(collection, priority)
	}
	return nil, nil
}

func (m *MockDatabase) UpdateTasks(result tasks_domain.Task, collection string) error {
	if m.UpdateTasksFunc != nil {
		return m.UpdateTasksFunc(result, collection)
	}
	return nil
}

func (m *MockDatabase) DeleteTasks(result tasks_domain.Task, collection string) error {
	if m.DeleteTasksFunc != nil {
		return m.DeleteTasksFunc(result, collection)
	}
	return nil
}

func (m *MockDatabase) TaskMigration(ctx context.Context) error {
	if m.TaskMigrationFunc != nil {
		return m.TaskMigrationFunc(ctx)
	}
	return nil
}
