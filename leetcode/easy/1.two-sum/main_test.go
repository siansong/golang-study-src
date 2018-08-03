package main

import (
	"testing"
)

type TestCase struct {
	Arr    []int
	Target int
	Result []int
}

var cases = []TestCase{
	{Arr: []int{1, 2, 3}, Target: 4, Result: []int{0, 2}},
	{Arr: []int{0, 9, 3}, Target: 3, Result: []int{0, 2}},
	{Arr: []int{1, 4, 3}, Target: 5, Result: []int{0, 1}},
	{Arr: []int{1, 4, 6, 3}, Target: 4, Result: []int{0, 3}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {

	validateFn := func(fn func(nums []int, target int) []int) {
		//normal test
		for _, val := range cases {
			ret := fn(val.Arr, val.Target)
			if ret[0] != val.Result[0] || ret[1] != val.Result[1] {
				t.Errorf("result should be [0, 1], got: %v", ret)
			}
		}

		//nil test
		rs := fn([]int{233, 666}, 15)
		if rs != nil {
			t.Errorf("result should be nil, got: %v", rs)
		}
	}

	for _, fn := range []func(nums []int, target int) []int{twoSumFn0, twoSumFn1, twoSumFn2} {
		validateFn(fn)
	}

}
