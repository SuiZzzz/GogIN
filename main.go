package main

import (
	"GoGin/router"
)

func main() {
	// 创建服务

	r1 := router.NewRouter()
	_ = r1.Run(":8083")
}
