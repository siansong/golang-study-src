package main

import (
	"fmt"
	"math"
)

/**
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.


*/

func reverseV0(x int) int {
	//极端值处理
	if x == 0 {
		return 0
	}

	//统一先转为正数处理
	positive := x > 0
	if !positive {
		x = -x
	}

	//将每位数逐一放进数组
	array := make([]int, 0)

	for x > 0 {
		array = append(array, x%10)
		x = x / 10
	}
	array = append(array, x)

	//对反转的数组进行转换
	noneZeroFound := false
	re := 0
	for _, v := range array {
		if !noneZeroFound {
			if v != 0 {
				noneZeroFound = true
				re = v
			} else {
				continue
			}
		} else {
			re = re*10 + v
		}
	}

	re = re / 10

	if re > 2147483647 || re < -2147483647 {
		re = 0
	}

	if positive {
		return re
	}
	return -re
}

func main() {
	fmt.Println(2 ^ 32 - 1) //^是异或运算，所以结果时33，二不是math.Pow(2, 32) -1
	fmt.Println(math.Pow(2, 3))
	v := math.Pow(2, 31) - 1
	fmt.Println(v)
	fmt.Printf("%f\n", v)
	fmt.Println(2147483651 > 4294967295)
}
