package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/tariq-ventura/Proyecto-go/internal/db/mongo"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

type Database interface {
	InsertTasks(result tasks_domain.Task, collection string) error
	SelectTasks(collection string) ([]tasks_domain.Task, error)
	UpdateTasks(result tasks_domain.Task, collection string) error
}

func NewDatabase(ctx context.Context) (Database, error) {
	dbType := os.Getenv("DB_CONTEXT")
	switch dbType {
	case "mongo":
		fmt.Print("Using mongo")
		return mongo.SetupMongo(ctx), nil
	default:
		return nil, errors.New("unsupported database backend")
	}
}
