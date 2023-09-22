package router

import (
	"GoGin/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	group := engine.Group("/user")
	{
		group.POST("/register", api.Register())
		group.POST("/login", api.Login())

		group.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"name":   "suqin",
				"status": 200,
			})
		})
	}
	return engine
}
