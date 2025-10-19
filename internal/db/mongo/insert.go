package mongo

import (
	"fmt"

	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func (m *MongoClient) InsertTasks(result tasks_domain.Task, collection string) error {
	insert, err := m.database.Collection(collection).InsertOne(m.ctx, result)

	if err != nil {
		logs.LogError("MongoDB insert error", map[string]interface{}{"error": err.Error()})
		return err
	}

	logs.LogInfo("MongoDB insert success", map[string]interface{}{"insertedID": fmt.Sprintf("%v", insert.InsertedID)})

	return nil
}
