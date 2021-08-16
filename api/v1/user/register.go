package user

import (
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/zhangshanwen/splie/code"
	"github.com/zhangshanwen/splie/initialize/db"
	"github.com/zhangshanwen/splie/initialize/service"
	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/internal/response"
	"github.com/zhangshanwen/splie/model"
)

func Register(c *service.Context) (resp service.Res) {
	p := param.Register{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	user := model.User{Mobile: p.Mobile}
	g := db.G
	g = g.Begin()
	defer func() {
		if resp.Err == nil {
			g.Commit()
		} else {
			g.Rollback()
		}
	}()
	if resp.Err = g.Where(&user).First(&user).Error; resp.Err != nil && resp.Err != gorm.ErrRecordNotFound {
		resp.ResCode = code.DbError
		return
	}
	if user.Id > 0 {
		resp.Err = errors.New("mobile is existed")
		resp.ResCode = code.MobileIsExist
		return
	}
	if resp.Err = copier.Copy(&user, &p); resp.Err != nil {
		return
	}
	if resp.Err = user.SetPassword(p.Password); resp.Err != nil {
		resp.ResCode = code.SetPasswordField
		return
	}
	if resp.Err = g.Create(&user).Error; resp.Err != nil {
		resp.ResCode = code.DbError
		return
	}
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &user); resp.Err != nil {
		resp.ResCode = code.CopyParamError
		return
	}
	resp.Data = r
	return
}
