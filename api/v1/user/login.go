package user

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/splie/code"
	"github.com/zhangshanwen/splie/initialize/db"
	"github.com/zhangshanwen/splie/initialize/service"
	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/internal/response"
	"github.com/zhangshanwen/splie/model"
	"github.com/zhangshanwen/splie/tools"
)

func Login(c *service.Context) (resp service.Res) {
	p := param.Login{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	user := model.User{Mobile: p.Mobile}
	g := db.G
	if resp.Err = g.Where(&user).First(&user).Error; resp.Err != nil {
		resp.ResCode = code.DbError
		return
	}
	if !user.CheckPassword(p.Password) {
		resp.ResCode = code.ActPWdError
	}
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &user); resp.Err != nil {
		resp.ResCode = code.CopyParamError
		return
	}
	var token string
	token, resp.Err = tools.CreateToken(user.Id)
	if resp.Err != nil {
		resp.ResCode = code.CreateTokenFailed
		return
	}
	r.Authorization = token
	resp.Data = r
	return
}
