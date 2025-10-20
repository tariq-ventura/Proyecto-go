package tasks_infra

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tariq-ventura/Proyecto-go/internal/db"
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
	"github.com/tariq-ventura/Proyecto-go/internal/validations"
)

func (th *TaskHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	var posts []tasks_domain.Task

	if err := c.BindJSON(&posts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	validate := validations.StructValidator

	done := make(chan string, len(posts))
	errors := make(chan error, len(posts))

	database, err := db.NewDatabase(ctx)
	if err != nil {
		logs.LogError("Database connection error", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal setup error"})
		return
	}

	for _, post := range posts {
		go func(post tasks_domain.Task) {
			response := validations.ValidateStruct(validate, post)
			if !response {
				errors <- fmt.Errorf("validation failed for task: %s", post.Name)
				return
			}

			dberr := database.DeleteTasks(post, "tasks")
			if dberr != nil {
				errors <- fmt.Errorf("database update error for task: %s, error: %v", post.Name, dberr)
				return
			}

			done <- post.Name
		}(post)
	}

	for i := 0; i < len(posts); i++ {
		select {
		case err := <-errors:
			logs.LogError("Task processing error", map[string]interface{}{"error": err.Error()})
		case <-done:
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "One or more tasks failed to process"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
}
