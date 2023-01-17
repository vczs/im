package help

import (
	"fmt"
	"runtime"
)

func VczsLog(desc string, err error) {
	p, _, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("error: [%s(%d)](%s):%v\n", name, line, desc, err)
}
