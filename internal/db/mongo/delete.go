package mongo

import (
	"fmt"

	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MongoClient) DeleteTasks(data tasks_domain.Task, collection string) error {
	objectID, err := primitive.ObjectIDFromHex(data.ID)
	if err != nil {
		logs.LogError("Invalid ObjectID", map[string]interface{}{"error": err.Error(), "id": data.ID})
		return err
	}

	filter := bson.M{"_id": objectID}

	result, err := m.database.Collection(collection).DeleteOne(m.ctx, filter)

	if err != nil {
		logs.LogError("MongoDB update error", map[string]interface{}{"error": err.Error()})
		return err
	}

	logs.LogInfo("MongoDB update success", map[string]interface{}{"Objects Deleted": fmt.Sprintf("%v", result.DeletedCount)})

	return nil
}
