package goutils

import (
	"fmt"
	"testing"
)

func TestParseUint(t *testing.T) {
	for i := uint64(0); i < 128; i++ {
		s := fmt.Sprintf("0b%b", i)
		v, err := ParseUint(s, 32)
		if err != nil {
			t.Error(err)
		}
		if v != i {
			t.Error(err)
		}
	}

	for i := uint64(0); i < 128; i++ {
		s := fmt.Sprintf("0B%b", i)
		v, err := ParseUint(s, 32)
		if err != nil {
			t.Error(err)
		}
		if v != i {
			t.Error(err)
		}
	}

	for i := uint64(0); i < 128; i++ {
		s := fmt.Sprintf("%d", i)
		v, err := ParseUint(s, 32)
		if err != nil {
			t.Error(err)
		}
		if v != i {
			t.Error(err)
		}
	}

	for i := uint64(0); i < 128; i++ {
		s := fmt.Sprintf("0x%x", i)
		v, err := ParseUint(s, 32)
		if err != nil {
			t.Error(err)
		}
		if v != i {
			t.Error(err)
		}
	}
}
