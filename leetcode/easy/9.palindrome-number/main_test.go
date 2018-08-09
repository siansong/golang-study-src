package main

import (
	"testing"
)

func TestPalidrome(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{233, false},
		{666, true},
		{21120, false},
	}

	for i, fn := range []func(int) bool{isPalindromeV0, isPalindromeV1, isPalindromeV2} {
		for _, c := range cases {
			r := fn(c.in)
			if (c.out && !r) || (!c.out && r) {
				t.Errorf("fn%v Failed on %v", i, c)
			}
		}
	}

}
