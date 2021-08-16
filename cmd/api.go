package main

import (
	"github.com/zhangshanwen/splie/initialize"
	"github.com/zhangshanwen/splie/router"
)

func main() {
	initialize.Initialize() // 注册服务
	router.InitRouter()     // 注册路由
}
