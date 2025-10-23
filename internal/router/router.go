package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Routes  *gin.Engine
	Context *gin.Context
}

func (r *Routes) SetupRouter() *gin.Engine {
	r.Routes = gin.Default()

	if _, ok := os.LookupEnv("TEST_ENV"); !ok {
		r.Routes.LoadHTMLGlob("templates/*")
	}

	r.SetupCors()

	r.Routes.Static("/static", "./static")
	r.Routes.GET("/healthz", r.Print)

	r.IndexRoutes(r.Routes)
	r.TasksRoutes(r.Routes)
	
	return r.Routes
}

func (r *Routes) Run() {
	r.Routes.Run(":3000")
}

func (r *Routes) Print(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to my API with Golang"})
}
