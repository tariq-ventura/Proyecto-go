package mongo

import (
	"context"
	"os"

	"github.com/tariq-ventura/Proyecto-go/internal/logs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client   *mongo.Client
	ctx      context.Context
	database *mongo.Database
}

func SetupMongo(ctx context.Context) *MongoClient {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))

	if err != nil {
		logs.LogError("MongoDB connection error", map[string]interface{}{"error": err.Error()})
		panic("app stopped")
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		logs.LogError("MongoDB ping error", map[string]interface{}{"error": err.Error()})
		panic("db conexion failed")
	}

	name := os.Getenv("DB_NAME")
	database := client.Database(name)

	return &MongoClient{
		client:   client,
		ctx:      ctx,
		database: database,
	}
}
