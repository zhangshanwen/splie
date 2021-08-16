package image

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/plug/client"
	pb "github.com/zhangshanwen/plug/proto"
	"github.com/zhangshanwen/splie/initialize/db"
	l "github.com/zhangshanwen/splie/initialize/logger"
	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/model"
	"io"
	"strings"
	"time"
)

func Pull(c *gin.Context) {
	p := param.ImagePull{}

	var err error
	if err = c.Bind(&p); err != nil {
		return
	}
	s := model.Server{}
	if err = db.G.First(&s, p.ServerId).Error; err != nil {
		return
	}
	var cli client.Client
	cli, err = client.NewKeep(context.Background(), time.Second*3, s.Address)
	if err != nil {
		return
	}
	defer cli.Close()
	pbc := pb.NewImageAliveClient(cli.Conn)
	stream, err := pbc.ImagePull(context.Background(), &pb.ImagePullRequest{RefStr: p.ImageName})
	if err != nil {
		l.Logger.Error(err)
		return
	}

	c.Stream(func(w io.Writer) bool {
		select {
		default:
			feature, err := stream.Recv()
			if err == io.EOF || err != nil {
				l.Logger.Error(err)
				c.SSEvent("message", "ok")
				return false
			}
			var data []map[string]interface{}
			for _, body := range strings.Split(string(feature.Body), "\r\n") {
				var bodyMap map[string]interface{}
				_ = json.Unmarshal([]byte(body), &bodyMap)
				data = append(data, bodyMap)
			}
			c.SSEvent("message", data)
			return true
		}
	})
}
