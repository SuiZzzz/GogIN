package serializer

const (
	OK = iota
	IllegalData
	UsernameIsExist
	UsernameNotExist
	IncorrectPassword
	IllegalPassword
	NicknameIsExist
	SystemError

	Succeed      = "成功"
	UserIsExit   = "用户已存在"
	UserNotExit  = "用户不存在"
	DataError    = "数据异常"
	ServerError  = "服务器数据异常"
	LoginSucceed = "登录成功"
)

type Response struct {
	Code    int
	Data    any
	Message string
	Error   error
}

type TokenData struct {
	User  any
	Token string
}
