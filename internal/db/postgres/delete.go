package postgres

import (
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) DeleteTasks(result tasks_domain.Task, collection string) error {
	err := db.client.Where("id = ?", result.ID).Delete(&tasks_domain.Task{}).Error

	if err != nil {
		logs.LogError("Error deleting task in Postgres", map[string]interface{}{"error": err.Error(), "task_id": result.ID})
		return err
	}

	logs.LogInfo("Successfully deleted task in Postgres", map[string]interface{}{"task_id": result.ID})
	return nil
}
