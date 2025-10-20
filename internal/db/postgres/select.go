package postgres

import (
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) SelectTasks(collection string) ([]tasks_domain.Task, error) {
	var results []tasks_domain.Task

	if err := db.client.Find(&results).Error; err != nil {
		logs.LogError("Error selecting tasks from Postgres", map[string]interface{}{"error": err.Error()})
		return nil, err
	}

	logs.LogInfo("Successfully selected tasks from Postgres", map[string]interface{}{"count": len(results)})
	return results, nil
}
