package tasks_infra

import (
	tasks_domain "github.com/tariq-ventura/Proyecto-go/internal/tasks/domain"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
}

func NewTaskHandler(server *gin.Context) tasks_domain.TaskInterface {
	return &TaskHandler{}
}
