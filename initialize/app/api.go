package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/splie/middleware"
)

var R = &gin.Engine{}

func init() {
	gin.ForceConsoleColor()
	R = gin.Default()
	R.Use(middleware.Cors())
}
