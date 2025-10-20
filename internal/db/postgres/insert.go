package postgres

import (
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (db *PostgresClient) InsertTasks(result tasks_domain.Task, collection string) error {
	insert := db.client.Create(&result)

	if insert.Error != nil {
		logs.LogError("PostgreSQL insert error", map[string]interface{}{"error": insert.Error.Error()})
		return insert.Error
	}

	logs.LogInfo("PostgreSQL insert success in collection"+collection, map[string]interface{}{"insertedID": result.ID})
	return nil
}
