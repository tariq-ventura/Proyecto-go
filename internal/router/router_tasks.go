package router

import (
	tasks_infra "github.com/tariq-ventura/Proyecto-go/internal/tasks/infra"

	"github.com/gin-gonic/gin"
)

func (ro *Routes) TasksRoutes(r *gin.Engine) {
	tr := tasks_infra.NewTaskHandler(ro.Context)

	routes := r.Group("/api/tasks")
	{
		routes.POST("/", tr.Insert)
		routes.GET("/", tr.Select)
		routes.PUT("/", tr.Update)
		routes.DELETE("/", tr.Delete)
	}
}
