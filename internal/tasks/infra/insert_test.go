package tasks_infra

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tariq-ventura/Proyecto-go/internal/db"
	"github.com/tariq-ventura/Proyecto-go/internal/mocks"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func TestInsert(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success - Insert a task", func(t *testing.T) {
		// Setup
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		task := []tasks_domain.Task{{
			Name:        "New Task",
			Description: "A description",
			Status:      "Pending",
			DueDate:     "2025-01-01",
			Priority:    "High",
		}}
		body, _ := json.Marshal(task)
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		mockDB := &mocks.MockDatabase{
			InsertTasksFunc: func(result tasks_domain.Task, collection string) error {
				return nil
			},
		}
		originalNewDatabase := db.NewDatabase
		overrideNewDatabase(mockDB, nil)
		defer func() { db.NewDatabase = originalNewDatabase }()

		handler := NewTaskHandler(nil)

		// Execute
		handler.Insert(c)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		var response gin.H
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Data created", response["message"])
	})

	t.Run("Error - Invalid payload", func(t *testing.T) {
		// Setup
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBufferString("invalid json"))
		c.Request.Header.Set("Content-Type", "application/json")

		handler := NewTaskHandler(nil)

		// Execute
		handler.Insert(c)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
