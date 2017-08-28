package goutils

import (
	"fmt"
	"io"
	"runtime"
)

func Trace(w io.Writer) {
	f, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fmt.Fprintf(w, "%s (%s:%d)\n", runtime.FuncForPC(f).Name(), file, line)
}
