package goutils

import (
	"bytes"
	"os"
	"testing"
)

func Test_trace(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	expect := "github.com/choueric/goutils.Test_trace (" + gopath + "/src/github.com/choueric/goutils/trace_test.go:13)\n"
	var result bytes.Buffer
	Trace(&result)

	if expect != result.String() {
		t.Error("Trace info dose not match.")
		t.Error(result.String())
		t.Error(expect)
	}
}
