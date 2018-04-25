package goutils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func PrefixStringPerLine(w io.Writer, lines, prefixStr string) {
	if lines == "" {
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		fmt.Fprintln(w, prefixStr+scanner.Text())
	}
}
