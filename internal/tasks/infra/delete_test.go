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

func TestDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewTaskHandler(nil)

	validTask := []tasks_domain.Task{{
		ID:          "1",
		Name:        "Valid Task to Delete",
		Description: "A valid description",
		Status:      "Pending",
		DueDate:     "2025-01-01",
		Priority:    "Low",
	}}

	t.Run("Success - Delete tasks", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(validTask) // Usar la tarea válida
		c.Request, _ = http.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(body))

		mockDB := &mocks.MockDatabase{
			DeleteTasksFunc: func(result tasks_domain.Task, collection string) error {
				return nil // Simula éxito
			},
		}
		originalNewDatabase := db.NewDatabase
		overrideNewDatabase(mockDB, nil)
		defer func() { db.NewDatabase = originalNewDatabase }()

		handler.Delete(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "Data deleted successfully"}`, w.Body.String())
	})

	t.Run("Error - Database error on delete", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(validTask) // Usar la tarea válida
		c.Request, _ = http.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(body))

		mockDB := &mocks.MockDatabase{
			DeleteTasksFunc: func(result tasks_domain.Task, collection string) error {
				return errors.New("db delete failed") // Simula error
			},
		}
		originalNewDatabase := db.NewDatabase
		overrideNewDatabase(mockDB, nil)
		defer func() { db.NewDatabase = originalNewDatabase }()

		handler.Delete(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"error": "One or more tasks failed to process"}`, w.Body.String())
	})
}
