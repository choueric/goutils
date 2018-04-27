package goutils

import (
	"bufio"
	"strings"
)

// @r can be get from *os.File by bufio.NewReader()
func ReadLine(r *bufio.Reader) (string, error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Trim(str, "\r\n")
	return str, nil
}
