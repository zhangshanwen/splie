package param

import "fmt"

type (
	Pagination struct {
		Page     int    `json:"page"      form:"page"`
		PageSize int    `json:"page_size" form:"page_size"`
		Sort     string `json:"sort"      form:"sort"`
		Asc      bool   `json:"asc"       form:"asc"`
	}
)

func (p *Pagination) Offset() int {
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) Order() string {
	if p.Sort == "" {
		p.Sort = "id"
	}
	if p.Asc {
		return fmt.Sprintf(" %s ", p.Sort)
	} else {
		return fmt.Sprintf(" %s DESC", p.Sort)
	}
}
