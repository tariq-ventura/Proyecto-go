package mongo

import (
	"context"

	"github.com/tariq-ventura/Proyecto-go/internal/logs"
)

func (m *MongoClient) TaskMigration(ctx context.Context) error {
	err := m.client.Ping(ctx, nil)

	if err != nil {
		logs.LogError("MongoDB ping error", map[string]interface{}{"error": err.Error()})
		panic("db conexion failed")
	}

	return nil
}
