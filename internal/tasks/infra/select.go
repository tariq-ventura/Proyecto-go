package tasks_infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tariq-ventura/Proyecto-go/internal/db"
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
)

func (th *TaskHandler) Select(c *gin.Context) {
	ctx := c.Request.Context()

	database, err := db.NewDatabase(ctx)
	if err != nil {
		logs.LogError("Database connection error", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal setup error"})
		return
	}

	tasks, dberr := database.SelectTasks("tasks")

	if dberr != nil {
		logs.LogError("Database Insert error", map[string]interface{}{"error": dberr.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal database insert error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}
