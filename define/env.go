package define

import "os"

var JwtKey = "vczs-key" // jwt key

var CodeLength = 6                // 验证码长度
var CodeExpire = 300              // 验证码过期时间（s）
var MainPwd = os.Getenv("163pwd") // 邮箱密码

var TokenExpire = 9600         // token过期时间
var RefreshTokenExpire = 96000 // refreshToken过期时间

var DefaultLimit int64 = 10 // 分页查询默认每页条数

var UserRoomTypeAlone = 1 // 私聊
var UserRoomTypeGroup = 2 // 群聊
