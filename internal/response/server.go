package response

import "github.com/zhangshanwen/splie/model"

type (
	ServerResponse struct {
		List       []model.Server `json:"list"`
		Pagination Pagination     `json:"pagination"`
	}
)
