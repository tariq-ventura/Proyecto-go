package tasks_infra

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/tariq-ventura/Proyecto-go/internal/db"
	"github.com/tariq-ventura/Proyecto-go/internal/logs"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
	"github.com/tariq-ventura/Proyecto-go/internal/validations"

	"github.com/gin-gonic/gin"
)

func (th *TaskHandler) Insert(c *gin.Context) {
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

	for i, post := range posts {
		go func(i int, post tasks_domain.Task) {
			namespace := uuid.New()
			name := []byte("tasks")
			id := uuid.NewSHA1(namespace, name)

			posts[i].ID = id.String()

			response := validations.ValidateStruct(validate, post)
			if !response {
				errors <- fmt.Errorf("validation failed for task: %s", post.Name)
				done <- post.Name
				return
			}

			dberr := database.InsertTasks(post, "tasks")
			if dberr != nil {
				errors <- fmt.Errorf("database insert error for task: %s, error: %v", post.Name, dberr)
				done <- post.Name
				return
			}

			done <- post.Name
		}(i, post)
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

	c.JSON(http.StatusOK, gin.H{"message": "Data created"})
}
