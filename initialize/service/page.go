package service

import (
	"gorm.io/gorm"

	"github.com/zhangshanwen/splie/internal/param"
	"github.com/zhangshanwen/splie/internal/response"
)

func GetPagination(db *gorm.DB, p *param.Pagination, r *response.Pagination, data interface{}) (err error) {
	if err = db.Count(&r.Total).Error; err != nil {
		return
	}
	if err = db.Order(p.Order()).Limit(p.PageSize).Offset(p.Offset()).Find(data).Error; err != nil {
		return
	}
	r.PageSize = p.PageSize
	r.Page = p.Page
	return
}
