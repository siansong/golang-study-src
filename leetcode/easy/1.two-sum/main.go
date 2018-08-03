package main

import (
	"fmt"
)

func hi() {
	fmt.Println("hi")
}

/**
https://leetcode.com/problems/two-sum/description/


Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

//v1,暴力遍历，时间复杂度：n^2, 空间复杂度k
func twoSumFn0(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//v2,2次遍历计算差值，检查存在，时间复杂度n, 空间复杂度n
func twoSumFn1(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, v := range nums {
		m[v] = idx
	}
	for idxX, x := range nums {
		idxY, ok := m[target-x]
		if ok && idxX != idxY {
			return []int{idxX, idxY}
		}
	}
	return nil
}

//v3,单次遍历计算差值，检查存在，时间复杂度n, 空间复杂度n (时间复杂度和v2一样，但是具体遍历次数确实少了1轮)
func twoSumFn2(nums []int, target int) []int {
	m := make(map[int]int)
	for indexY, y := range nums {
		indexX, ok := m[target-y]
		if ok && indexY != indexX {
			return []int{indexX, indexY}
		}
		m[y] = indexY
	}
	return nil
}

//TODO goroutine 优化
//TODO 算法理论优化

func main() {
	rs := twoSumFn0([]int{2, 7, 11, 15}, 9)
	fmt.Println(rs)
}
