package user

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/splie/code"
	"github.com/zhangshanwen/splie/initialize/service"
	"github.com/zhangshanwen/splie/internal/response"
)

func Detail(c *service.Context) (resp service.Res) {
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &c.User); resp.Err != nil {
		resp.ResCode = code.CopyParamError
		return
	}
	resp.Data = r
	return
}
