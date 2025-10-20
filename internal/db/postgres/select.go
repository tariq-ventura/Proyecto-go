package postgres

import (
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) SelectTasks(collection string) ([]tasks_domain.Task, error) {
	var results []tasks_domain.Task

	if err := db.client.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
