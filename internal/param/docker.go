package param

type (
	Images struct {
		ServerId int `json:"server_id"  form:"server_id" binding:"required"`
	}
	DockerIdParam struct {
		ServerId int `uri:"server_id"           binding:"required"`
	}
	ImageSearch struct {
		ServerId  int    `json:"server_id"  form:"server_id" binding:"required"`
		ImageName string `json:"image_name"  form:"image_name" binding:"required"`
	}

	ImagePull struct {
		ServerId  int    `json:"server_id"  form:"server_id" binding:"required"`
		ImageName string `json:"image_name"  form:"image_name" binding:"required"`
	}
)
