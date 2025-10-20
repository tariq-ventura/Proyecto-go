package postgres

import (
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) SelectTasksStatus(collection, status string) ([]tasks_domain.Task, error) {
	var results []tasks_domain.Task

	if err := db.client.Where("status = ?", status).Find(&results).Error; err != nil {
		logs.LogError("Error selecting tasks by status from Postgres", map[string]interface{}{"status": status, "error": err.Error()})
		return nil, err
	}

	logs.LogInfo("Successfully selected tasks by status from Postgres", map[string]interface{}{"status": status, "count": len(results)})
	return results, nil
}
