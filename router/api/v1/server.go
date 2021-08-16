package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/splie/api/v1/server"
)

func InitServerRouter(Router *gin.RouterGroup) {
	r := Router.Group("servers")
	{
		r.GET("", v(server.Servers))       // 获取服务列表
		r.POST("", v(server.Register))     // 服务注册
		r.PUT("/:id", v(server.Update))    // 更新服务
		r.DELETE("/:id", v(server.Delete)) // 注销服务
	}
}
