package postgres

import (
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) SelectTasksDate(collection, date string) ([]tasks_domain.Task, error) {
	var results []tasks_domain.Task

	if err := db.client.Where("dueDate = ?", date).Find(&results).Error; err != nil {
		logs.LogError("Error selecting tasks by date from Postgres", map[string]interface{}{"date": date, "error": err.Error()})
		return nil, err
	}

	logs.LogInfo("Successfully selected tasks by dueDate from Postgres", map[string]interface{}{"date": date, "count": len(results)})
	return results, nil
}
