package mongo

import (
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoClient) SelectTasks(collection string) ([]tasks_domain.Task, error) {
	var results []tasks_domain.Task

	cur, err := m.database.Collection(collection).Find(m.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(m.ctx)

	for cur.Next(m.ctx) {
		var result tasks_domain.Task
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
