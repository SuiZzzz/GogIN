package api

import (
	"GoGin/service"
	"GoGin/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Talk() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := types.AIReq{}
		_ = ctx.ShouldBind(&req)
		sparkService := service.GetSparkServiceInstance()
		resp := sparkService.ConnToSpark(&req)
		ctx.JSON(http.StatusOK, resp)
	}
}
