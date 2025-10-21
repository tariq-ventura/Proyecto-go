package tasks_infra

import (
	"context"
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

func overrideNewDatabase(mockDB db.Database, err error) {
	db.NewDatabase = func(ctx context.Context) (db.Database, error) {
		return mockDB, err
	}
}

func TestSelect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success - Select tasks", func(t *testing.T) {
		// Setup
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/api/tasks", nil)

		mockTasks := []tasks_domain.Task{{ID: "1", Name: "Test Task"}}
		mockDB := &mocks.MockDatabase{
			SelectTasksFunc: func(collection string) ([]tasks_domain.Task, error) {
				return mockTasks, nil
			},
		}
		originalNewDatabase := db.NewDatabase
		overrideNewDatabase(mockDB, nil)
		defer func() { db.NewDatabase = originalNewDatabase }()

		handler := NewTaskHandler(nil)

		// Execute
		handler.Select(c)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		var response gin.H
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.NotNil(t, response["data"])
	})

	t.Run("Error - Database error", func(t *testing.T) {
		// Setup
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/api/tasks", nil)

		mockDB := &mocks.MockDatabase{
			SelectTasksFunc: func(collection string) ([]tasks_domain.Task, error) {
				return nil, errors.New("db error")
			},
		}
		originalNewDatabase := db.NewDatabase
		overrideNewDatabase(mockDB, nil)
		defer func() { db.NewDatabase = originalNewDatabase }()

		handler := NewTaskHandler(nil)

		// Execute
		handler.Select(c)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
