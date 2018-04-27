package goutils

import "fmt"

func PrintStringMismatch(result, original []byte, debug bool) {
	if len(result) != len(original) {
		fmt.Println("length doesn't match:", len(result), len(original))
		if debug {
			fmt.Println("============== result =============")
			fmt.Println(string(result))
		}
		return
	}
	for i := 0; i < len(result) && i < len(original); i++ {
		if result[i] != original[i] {
			fmt.Println("Content doesn't match, index:", i)
			fmt.Printf("result: %s(%d), original: %s(%d)\n",
				string(result[i]), result[i],
				string(original[i]), original[i])
			if debug {
				fmt.Println("============== result =============")
				fmt.Println(string(result))
			}
			break
		}
	}
}
