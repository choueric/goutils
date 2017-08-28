package goutils

import (
	"fmt"
	"runtime"
)

func Trace() {
	f, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fmt.Printf("%s (%s:%d)\n", runtime.FuncForPC(f).Name(), file, line)
}
