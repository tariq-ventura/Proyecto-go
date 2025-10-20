package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ro *Routes) IndexRoutes(r *gin.Engine) {
	r.GET("/", indexHandler)
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
