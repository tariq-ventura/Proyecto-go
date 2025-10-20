package tasks_domain

import "github.com/gin-gonic/gin"

type TaskInterface interface {
	Insert(c *gin.Context)
	Select(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
