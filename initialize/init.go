package initialize

import (
	"github.com/zhangshanwen/splie/initialize/conf"
	"github.com/zhangshanwen/splie/initialize/db"
	"github.com/zhangshanwen/splie/initialize/logger"
)

func Initialize() {
	logger.InitLog() // 初始化日志
	conf.InitConf()  // 初始化配置文件
	db.InitDb()      // 初始化数据库
}
