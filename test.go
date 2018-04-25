package goutils

import "fmt"

func PrintStringMismatch(a, b []byte, debug bool) {
	if len(a) != len(b) {
		fmt.Println("length doesn't match:", len(a), len(b))
		if debug {
			fmt.Println(string(a), "----------------------\n", string(b))
		}
		return
	}
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			fmt.Println(i, string(a[i]), string(b[i]))
			break
		}
	}
}
