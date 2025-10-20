package postgres

import (
	"context"

	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) TaskMigration(ctx context.Context) error {
	err := db.client.AutoMigrate(tasks_domain.Task{})

	if err != nil {
		logs.LogError("Error during task migration in Postgres", map[string]interface{}{"error": err.Error()})
		return err
	}

	logs.LogInfo("Successfully completed task migration in Postgres", nil)
	return nil
}
