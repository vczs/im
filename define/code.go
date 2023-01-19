package define

const (
	OK = 0 // success

	TOKEN_INVALID             = 101 // token无效
	PARAMETER_ANAIYSIS_FAILED = 102 // 参数解析失败
	REQUEST_OFTEN             = 103 // 请求频繁
	ACCESS_DENIED             = 104 // 拒绝访问
	PARAMETER_WRONG           = 105 // 参数有误

	EMAIL_EMPTY               = 10001 // 邮箱为空
	EMAIL_HAS_REGISTERED      = 10002 // 邮箱已被注册
	ACCOUNT_OR_PASSWORD_EMPTY = 10010 // 账号或密码为空
	ACCOUNT_OR_PASSWORD_ERROR = 10011 // 账号或密码错误
	ACCOUNT_EXIST             = 10012 // 账号已存在
	EMAIL_CODE_WRONG          = 10013 // 验证码不正确
	USER_NOT_EXIST            = 10014 // 用户不存在

	NOT_ADD_YOURSELF = 20001 // 不能添加自己
	ALREADY_FRIEND   = 20002 // 你与该用户已是好友
)

var message = map[int]string{
	OK: "success",

	TOKEN_INVALID:             "token无效!",
	PARAMETER_ANAIYSIS_FAILED: "参数解析失败!",
	REQUEST_OFTEN:             "请求频繁!",
	ACCESS_DENIED:             "拒绝访问!",
	PARAMETER_WRONG:           "参数有误!",

	EMAIL_EMPTY:               "邮箱为空!",
	EMAIL_HAS_REGISTERED:      "邮箱已被注册!",
	ACCOUNT_OR_PASSWORD_EMPTY: "账号或密码为空!",
	ACCOUNT_OR_PASSWORD_ERROR: "账号或密码错误!",
	ACCOUNT_EXIST:             "账号已存在!",
	EMAIL_CODE_WRONG:          "验证码不正确!",
	USER_NOT_EXIST:            "用户不存在!",

	NOT_ADD_YOURSELF: "不能添加自己!",
	ALREADY_FRIEND:   "你与该用户已是好友!",
}

// GetMessage 获取message
func Message(code int) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "服务器发生未知错误~"
	}
}
