package server

import (
	"github.com/zhangshanwen/splie/code"
	"github.com/zhangshanwen/splie/initialize/db"
	"github.com/zhangshanwen/splie/initialize/service"
	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/model"
)

func Delete(c *service.Context) (resp service.Res) {
	idp := param.ServersIdParam{}
	if resp.Err = c.ShouldBindUri(&idp); resp.Err != nil {
		resp.ResCode = code.IdError
		return
	}
	s := model.Server{}
	if resp.Err = db.G.First(&s, idp.Id).Error; resp.Err != nil {
		resp.ResCode = code.IdError
		return
	}
	s.IsDeleted = true
	if resp.Err = db.G.Save(&s).Error; resp.Err != nil {
		resp.ResCode = code.DbError
		return
	}
	resp.Data = s
	return
}
