package api

import (
	"GoGin/service"
	"GoGin/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 注册用户
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRequest *types.UserRegisterReq
		userService := service.GetUserServiceInstance()
		_ = ctx.ShouldBind(&userRequest)
		res := userService.Register(userRequest, ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	}
}

// Login 用户登录
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginRequest *types.UserLoginReq
		userService := service.GetUserServiceInstance()
		_ = ctx.ShouldBind(&loginRequest)
		res := userService.Login(loginRequest, ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	}
}
