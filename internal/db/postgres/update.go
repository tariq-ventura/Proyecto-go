package postgres

import (
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) UpdateTasks(result tasks_domain.Task, collection string) error {
	err := db.client.Save(&tasks_domain.Task{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Status:      result.Status,
		DueDate:     result.DueDate,
		Priority:    result.Priority,
	}).Error

	if err != nil {
		logs.LogError("Error updating task in Postgres", map[string]interface{}{"error": err.Error(), "task_id": result.ID})
		return err
	}

	logs.LogInfo("Successfully updated task in Postgres", map[string]interface{}{"task_id": result.ID})

	return nil
}
