package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Routes  *gin.Engine
	Context *gin.Context
}

func (r *Routes) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/healthz", r.Print)

	r.IndexRoutes(router)
	r.TasksRoutes(router)

	return router
}

func (r *Routes) Run() {
	r.Routes.Run(":3000")
}

func (r *Routes) Print(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to my API with Golang"})
}
