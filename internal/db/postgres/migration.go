package postgres

import (
	"context"

	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) TaskMigration(ctx context.Context) error {
	err := db.client.AutoMigrate(tasks_domain.Task{})

	if err != nil {
		return err
	}
	return nil
}
