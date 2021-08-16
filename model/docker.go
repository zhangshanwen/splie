package model

type (
	Images struct {
		BaseModel
		NickName string `json:"nick_name"`
		ImageId  string `json:"image_id"`
		ServerId int    `json:"server_id"`
		Intro    string `json:"intro"`
	}
)
