package server

import (
	"gorm.io/gorm"

	"github.com/zhangshanwen/splie/code"
	"github.com/zhangshanwen/splie/initialize/db"
	"github.com/zhangshanwen/splie/initialize/service"
	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/internal/response"
	"github.com/zhangshanwen/splie/model"
)

func Servers(c *service.Context) (resp service.Res) {
	p := param.ServersParam{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	r := response.ServerResponse{}
	g := db.G.Model(&model.Server{})
	if resp.Err = service.GetPagination(g, &p.Pagination, &r.Pagination, &r.List); resp.Err != nil && resp.Err != gorm.ErrRecordNotFound {
		resp.ResCode = code.DbError
		return
	}
	resp.Data = r
	return
}
