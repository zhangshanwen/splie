package server

import (
	"time"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/plug/client"
	pb "github.com/zhangshanwen/plug/proto"
	"github.com/zhangshanwen/splie/code"
	"github.com/zhangshanwen/splie/initialize/db"
	"github.com/zhangshanwen/splie/initialize/service"
	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/model"
)

func Register(c *service.Context) (resp service.Res) {
	p := param.ServerRegisterParam{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	g := db.G.Begin()
	defer func() {
		if resp.Err != nil {
			g.Rollback()
		} else {
			g.Commit()
		}
	}()
	s := &model.Server{}
	if resp.Err = copier.Copy(&s, &p); resp.Err != nil {
		resp.ResCode = code.CopyParamError
		return
	}
	var cli client.Client
	cli, resp.Err = client.New(c, time.Second*3, s.Address)
	if resp.Err != nil {
		resp.ResCode = code.ConnServer
		return
	}
	defer cli.Close()
	pbc := pb.NewDockerClient(cli.Conn)
	_, resp.Err = pbc.Ping(cli.Ctx, &pb.PingRequest{})
	if resp.Err != nil {
		resp.ResCode = code.ConnServer
		return
	}
	// 心跳检测成功
	s.Ping = true
	if resp.Err = g.Create(&s).Error; resp.Err != nil {
		resp.ResCode = code.DbError
		return
	}
	return
}
