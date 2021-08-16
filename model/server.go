package model

type Server struct {
	BaseModel
	Address  string `json:"address"     gorm:"index:unique"`
	NickName string `json:"nick_name"`
	Intro    string `json:"intro"`
	Ping     bool   `json:"ping"` //心跳检测结果
}
