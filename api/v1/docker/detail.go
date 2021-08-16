package docker

import (
	"encoding/json"
	"time"

	"github.com/zhangshanwen/plug/client"
	pb "github.com/zhangshanwen/plug/proto"
	"github.com/zhangshanwen/splie/code"
	"github.com/zhangshanwen/splie/initialize/db"
	"github.com/zhangshanwen/splie/initialize/service"
	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/model"
)

func Detail(c *service.Context) (resp service.Res) {
	idp := param.DockerIdParam{}
	if resp.Err = c.ShouldBindUri(&idp); resp.Err != nil {
		resp.ResCode = code.IdError
		return
	}
	s := model.Server{}
	if resp.Err = db.G.First(&s, idp.ServerId).Error; resp.Err != nil {
		resp.ResCode = code.IdError
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
	var res *pb.InfoReply
	res, resp.Err = pbc.Info(cli.Ctx, &pb.InfoRequest{})
	if resp.Err != nil {
		resp.ResCode = code.ConnServer
		return
	}
	var data map[string]interface{}
	if resp.Err = json.Unmarshal(res.Info, &data); resp.Err != nil {
		resp.ResCode = code.JsonUnmarshalFailed
		return
	}
	resp.Data = data
	return
}
