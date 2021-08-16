package api

import (
	"github.com/zhangshanwen/splie/initialize/app"
	"github.com/zhangshanwen/splie/router/api/v1"
)

func RegisterApiV1Router() {
	api := app.R.Group("api")
	group := api.Group("v1")

	v1.InitUserRouter(group)
	v1.InitDockerRouter(group)
	v1.InitServerRouter(group)
}
