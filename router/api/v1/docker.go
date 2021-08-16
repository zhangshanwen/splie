package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/splie/api/v1/docker/image"

	"github.com/zhangshanwen/splie/api/v1/docker"
)

func InitDockerRouter(Router *gin.RouterGroup) {
	g1 := Router.Group("docker")
	{
		g1.GET("/:server_id", v(docker.Detail)) // 获取docker服务详情
	}
	g2 := g1.Group("images")
	{
		g2.GET("", v(image.Images))              // 获取镜像列表
		g2.GET("init", v(image.InitLocalImages)) // 初始化本地镜像
		g2.GET("search", v(image.Search))        // 搜索远程镜像
		g2.GET("pull", image.Pull)               // 搜索远程镜像
	}
}
