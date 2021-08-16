package code

/*
失败code 200000 开头
(200000,201000] -> 基本错误
(201000,202000] -> 认证错误
(202000,203000] -> 注册错误
(203000,204000] -> 登录错误

*/
const (
	BaseFailedCode      = 200000 // 基数失败code
	Failed              = 200001 // 失败
	ParamsError         = 200002 // 参数错误
	DbError             = 200003 // 数据库错误
	CopyParamError      = 200004 // 拷贝参数错误
	IdError             = 200005 // 错误id
	AuthFailed          = 201001 // 认证失败
	AuthInvalid         = 201002 // 认证失效
	MobileIsExist       = 202001 // 手机已经存在
	SetPasswordField    = 202002 // 设置密码错误
	LoginFailed         = 203001 // 登录失败
	CreateTokenFailed   = 203002 // 生成token失败
	ActPWdError         = 203003 // 账号/密码错误
	ConnServer          = 204001 // 连接服务错误
	JsonUnmarshalFailed = 204002 // json反序列化错误
)
