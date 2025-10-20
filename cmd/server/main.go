package main

import (
	"context"

	"github.com/tariq-ventura/Proyecto-go/internal/db"
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	"github.com/tariq-ventura/Proyecto-go/internal/router"
)

func main() {
	ctx := context.Background()

	r := &router.Routes{}
	r.Routes = r.SetupRouter()
	r.Run()

	database, err := db.NewDatabase(ctx)
	if err != nil {
		logs.LogError("Database connection error", map[string]interface{}{"error": err.Error()})
		panic("app stopped")
	}

	err = database.TaskMigration(ctx)
	if err != nil {
		logs.LogError("Database migration error", map[string]interface{}{"error": err.Error()})
		panic("app stopped")
	}
}
