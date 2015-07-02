package lm

import (
	"fmt"
	"path"
	"runtime"
)

func lookupCaller(extraCalldepth int) string {
	caller := ""
	_, file, line, ok := runtime.Caller(2 + extraCalldepth)
	if !ok {
		return "unkown:0"
	}
	caller = fmt.Sprintf("%s:%d", path.Base(file), line)
	return caller
}
