package router

import (
	"GoGin/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	group1 := engine.Group("/user")
	{
		group1.POST("/register", api.Register())
		group1.POST("/login", api.Login())
	}
	group2 := engine.Group("/ai")
	{
		group2.POST("/talk", api.Talk())
	}
	return engine
}
