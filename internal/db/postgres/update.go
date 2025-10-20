package postgres

import (
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
		return err
	}

	return nil
}
