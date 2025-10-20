package postgres

import tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"

func (db *PostgresClient) DeleteTasks(result tasks_domain.Task, collection string) error {
	err := db.client.Delete(&tasks_domain.Task{}, result.ID).Error

	if err != nil {
		return err
	}

	return nil
}
