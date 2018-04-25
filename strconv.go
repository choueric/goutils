package goutils

import (
	"regexp"
	"strconv"
)

// ParseUint like strconv.ParseUnit, but it recognize binary string like "0b0101"
// and parse it. 10 and 16 base is OK.
func ParseUint(str string, bitSize int) (uint64, error) {
	matched, err := regexp.MatchString(`^0[Bb][01]*`, str)
	if err != nil {
		return 0, err
	}

	if matched {
		return strconv.ParseUint(str[2:], 2, bitSize)
	} else {
		return strconv.ParseUint(str, 0, bitSize)
	}
}
