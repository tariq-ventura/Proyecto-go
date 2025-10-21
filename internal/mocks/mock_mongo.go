package mocks

import (
	"context"

	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

type MockMongoClient struct{}

func (m *MockMongoClient) InsertTasks(result tasks_domain.Task, collection string) error {
	return nil
}

func (m *MockMongoClient) SelectTasks(collection string) ([]tasks_domain.Task, error) {
	return nil, nil
}

func (m *MockMongoClient) SelectTasksStatus(collection, status string) ([]tasks_domain.Task, error) {
	return nil, nil
}

func (m *MockMongoClient) SelectTasksDate(collection, date string) ([]tasks_domain.Task, error) {
	return nil, nil
}

func (m *MockMongoClient) SelectTasksPriority(collection, priority string) ([]tasks_domain.Task, error) {
	return nil, nil
}

func (m *MockMongoClient) UpdateTasks(result tasks_domain.Task, collection string) error {
	return nil
}

func (m *MockMongoClient) DeleteTasks(result tasks_domain.Task, collection string) error {
	return nil
}

func (m *MockMongoClient) TaskMigration(ctx context.Context) error {
	return nil
}
