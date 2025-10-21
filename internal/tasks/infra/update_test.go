package tasks_infra

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tariq-ventura/Proyecto-go/internal/db"
	"github.com/tariq-ventura/Proyecto-go/internal/mocks"
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"
)

func TestUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewTaskHandler(nil)

	t.Run("Success - Update tasks", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		tasks := []tasks_domain.Task{{
			ID:          "1",
			Name:        "Valid Task Name",
			Description: "Valid Description",
			Status:      "Pending",
			DueDate:     "2025-01-01",
			Priority:    "Medium",
		}}
		body, _ := json.Marshal(tasks)
		c.Request, _ = http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))

		mockDB := &mocks.MockDatabase{
			UpdateTasksFunc: func(result tasks_domain.Task, collection string) error {
				return nil
			},
		}
		originalNewDatabase := db.NewDatabase
		overrideNewDatabase(mockDB, nil)
		defer func() { db.NewDatabase = originalNewDatabase }()

		handler.Update(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "Data updated successfully"}`, w.Body.String())
	})

	t.Run("Error - Database error on update", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		tasks := []tasks_domain.Task{{
			ID:          "1",
			Name:        "Valid Task Name",
			Description: "Valid Description",
			Status:      "Pending",
			DueDate:     "2025-01-01",
			Priority:    "Medium",
		}}
		body, _ := json.Marshal(tasks)
		c.Request, _ = http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))

		mockDB := &mocks.MockDatabase{
			UpdateTasksFunc: func(result tasks_domain.Task, collection string) error {
				return errors.New("db update failed") // Simula error
			},
		}
		originalNewDatabase := db.NewDatabase
		overrideNewDatabase(mockDB, nil)
		defer func() { db.NewDatabase = originalNewDatabase }()

		handler.Update(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "One or more tasks failed to process"}`, w.Body.String())
	})

	t.Run("Error - Invalid request payload", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPut, "/", bytes.NewBufferString(`[{"id":"1"`)) // JSON inválido

		handler.Update(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error": "Invalid request payload"}`, w.Body.String())
	})
}
