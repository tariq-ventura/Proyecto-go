package tasks_infra

import (
	"fmt"
	"net/http"
	"sync"

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

	database, err := db.NewDatabase(ctx)
	if err != nil {
		logs.LogError("Database connection error", map[string]interface{}{"error": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal setup error"})
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var processingErrors []error

	for _, post := range posts {
		wg.Add(1)
		go func(post tasks_domain.Task) {
			defer wg.Done()

			response := validations.ValidateStruct(validate, post)
			if !response {
				err := fmt.Errorf("validation failed for task: %s", post.Name)
				mu.Lock()
				processingErrors = append(processingErrors, err)
				mu.Unlock()
				return
			}

			dberr := database.DeleteTasks(post, "tasks")
			if dberr != nil {
				err := fmt.Errorf("database delete error for task: %s, error: %v", post.Name, dberr)
				mu.Lock()
				processingErrors = append(processingErrors, err)
				mu.Unlock()
				return
			}
		}(post)
	}

	wg.Wait()

	if len(processingErrors) > 0 {
		for _, e := range processingErrors {
			logs.LogError("Task processing error", map[string]interface{}{"error": e.Error()})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "One or more tasks failed to process"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
}
