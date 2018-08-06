package main

import (
	"testing"
)

type TestCase struct {
	num    int
	target int
}

var cases = []TestCase{
	{123, 321},
	{-123, -321},
	{120, 21},
	{1534236469, 0},
	{1563847412, 0},
}

func TestReverse(t *testing.T) {

	fns := []func(int) int{reverseV0}

	for _, fn := range fns {
		for _, v := range cases {
			rs := fn(v.num)
			if rs != v.target {
				t.Errorf("Failed [%v]  actual: %v", v, rs)
			}
		}
	}

}
