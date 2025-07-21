package router

import (
	"net/http"

	"github.com/3xhvy/go-backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", pongGet)
		v1.POST("/ping", pongPost)
		v1.PUT("/ping", pongPut)
		v1.DELETE("/ping", pongDelete)

		v1.GET("/user", controller.NewUserController().GetUserById)
	}

	return router
}

func pongGet(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong " + name,
		"uid":     uid,
	})
}

func pongPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong",
	})
}

func pongPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong",
	})
}

func pongDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong",
	})
}
