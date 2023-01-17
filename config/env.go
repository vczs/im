package config

import "os"

var JwtKey = "vczs-key" // jwt key

var CodeLength = 6                // 验证码长度
var CodeExpire = 300              // 验证码过期时间（s）
var MainPwd = os.Getenv("163pwd") // 邮箱密码

var TokenExpire = 9600         // token过期时间
var RefreshTokenExpire = 96000 // refreshToken过期时间
