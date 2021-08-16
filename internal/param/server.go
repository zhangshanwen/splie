package param

type (
	ServersParam struct {
		Pagination
	}
	ServerRegisterParam struct {
		Address  string `json:"address"     binding:"required"`
		NickName string `json:"nick_name"`
		Intro    string `json:"intro"`
	}
	ServersIdParam struct {
		Id int `uri:"id"           binding:"required"`
	}
	ServersUpdateParam struct {
		Address  string `json:"address"     binding:"required"`
		NickName string `json:"nick_name"`
		Intro    string `json:"intro"`
	}
)
