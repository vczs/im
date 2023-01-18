package define

const (
	OK = 0 // success

	TOKEN_INVALID             = 101 // token无效
	PARAMETER_ANAIYSIS_FAILED = 102 // 参数解析失败
	REQUEST_OFTEN             = 103 // 请求频繁
	ACCESS_DENIED             = 104 // 拒绝访问

	EMAIL_EMPTY               = 10001 // 邮箱为空
	EMAIL_HAS_REGISTERED      = 10002 // 邮箱已被注册
	ACCOUNT_OR_PASSWORD_EMPTY = 10010 // 账号或密码为空
	ACCOUNT_OR_PASSWORD_ERROR = 10011 // 账号或密码错误
)

var message = map[int]string{
	OK: "success",

	TOKEN_INVALID:             "token无效!",
	PARAMETER_ANAIYSIS_FAILED: "参数解析失败!",
	REQUEST_OFTEN:             "请求频繁!",
	ACCESS_DENIED:             "拒绝访问!",

	EMAIL_EMPTY:               "邮箱为空!",
	EMAIL_HAS_REGISTERED:      "邮箱已被注册!",
	ACCOUNT_OR_PASSWORD_EMPTY: "账号或密码为空!",
	ACCOUNT_OR_PASSWORD_ERROR: "账号或密码错误!",
}

// GetMessage 获取message
func Message(code int) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "服务器发生未知错误~"
	}
}
