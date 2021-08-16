package response

type (
	Pagination struct {
		Page     int   `json:"page"      form:"page"`
		PageSize int   `json:"page_size" form:"page_size"`
		Total    int64 `json:"total"     form:"total"`
	}
)
