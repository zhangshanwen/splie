package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/splie/api/v1/user"
)

func InitUserRouter(Router *gin.RouterGroup) {
	r := Router.Group("user")
	{
		r.POST("", v(user.Register))   // 创建用户
		r.POST("login", v(user.Login)) // 登录用户
		r.GET("", jwt(user.Detail))    // 获取用户信息
	}
}
