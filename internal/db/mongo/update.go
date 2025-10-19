package mongo

import (
	"fmt"

	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MongoClient) UpdateTasks(data tasks_domain.Task, collection string) error {
	objectID, err := primitive.ObjectIDFromHex(data.ID)
	if err != nil {
		logs.LogError("Invalid ObjectID", map[string]interface{}{"error": err.Error(), "id": data.ID})
		return err
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{"$set": bson.M{
		"name":        data.Name,
		"description": data.Description,
		"status":      data.Status,
		"dueDate":     data.DueDate,
		"priority":    data.Priority,
	}}

	result, err := m.database.Collection(collection).UpdateOne(m.ctx, filter, update)

	if err != nil {
		logs.LogError("MongoDB update error", map[string]interface{}{"error": err.Error()})
		return err
	}

	logs.LogInfo("MongoDB update success", map[string]interface{}{"UpsertedID": fmt.Sprintf("%v", result.UpsertedID)})

	return nil
}
