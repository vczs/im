package help

import (
	"im/define"
	"math/rand"
	"time"
)

func GenerateEmailCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(10)])
	}
	return code
}
