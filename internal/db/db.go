package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/tariq-ventura/Proyecto-go/internal/db/mongo"
	"github.com/tariq-ventura/Proyecto-go/internal/db/postgres"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

type Database interface {
	InsertTasks(result tasks_domain.Task, collection string) error
	SelectTasks(collection string) ([]tasks_domain.Task, error)
	SelectTasksStatus(collection, status string) ([]tasks_domain.Task, error)
	SelectTasksDate(collection, date string) ([]tasks_domain.Task, error)
	SelectTasksPriority(collection, priority string) ([]tasks_domain.Task, error)
	UpdateTasks(result tasks_domain.Task, collection string) error
	DeleteTasks(result tasks_domain.Task, collection string) error
	TaskMigration(ctx context.Context) error
}

var NewDatabase = func(ctx context.Context) (Database, error) {
	dbType := os.Getenv("DB_CONTEXT")
	switch dbType {
	case "mongo":
		fmt.Print("Using mongo")
		return mongo.SetupMongo(ctx), nil
	case "postgres":
		fmt.Print("Using postgres")
		return postgres.SetupPostgres(ctx), nil
	default:
		return nil, errors.New("unsupported database backend")
	}
}
