package router

import (
	"fmt"

	"github.com/zhangshanwen/splie/initialize/app"
	"github.com/zhangshanwen/splie/initialize/conf"
	"github.com/zhangshanwen/splie/router/api"
)

func InitRouter() {
	api.RegisterApiV1Router()
	run()
}

func run() {
	_ = app.R.Run(fmt.Sprintf(":%s", conf.C.Port))

}
