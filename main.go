package main

import (
	"GoGin/conf"
	"GoGin/router"
)

func main() {
	// 创建服务

	r1 := router.NewRouter()
	_ = r1.Run(":" + conf.Conf.Application.Port)
}
