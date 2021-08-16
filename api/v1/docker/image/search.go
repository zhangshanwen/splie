package image

import (
	"time"

	"github.com/zhangshanwen/plug/client"
	pb "github.com/zhangshanwen/plug/proto"
	"github.com/zhangshanwen/splie/code"
	"github.com/zhangshanwen/splie/initialize/db"
	"github.com/zhangshanwen/splie/initialize/service"
	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/model"
)

func Search(c *service.Context) (resp service.Res) {
	p := param.ImageSearch{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	s := model.Server{}
	if resp.Err = db.G.First(&s, p.ServerId).Error; resp.Err != nil {
		resp.ResCode = code.DbError
		return
	}
	var cli client.Client
	cli, resp.Err = client.New(c, time.Second*3, s.Address)
	if resp.Err != nil {
		resp.ResCode = code.ConnServer
		return
	}
	defer cli.Close()
	pbc := pb.NewImageClient(cli.Conn)
	resp.Data, resp.Err = pbc.ImageSearch(cli.Ctx, &pb.ImageSearchRequest{Term: p.ImageName, Limit: 100})
	if resp.Err != nil {
		resp.ResCode = code.ConnServer
		return
	}
	return
}
