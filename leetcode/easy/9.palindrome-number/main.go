package main

import (
	"fmt"
)

/**

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Follow up:

Coud you solve it without converting the integer to a string?
*/

//假设整数位数为n，则复杂度应该是3/2n左右
//思路是把整数拆到数组里，头尾相比
func isPalindromeV0(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	//split int to array
	arr := make([]int, 0)
	for x >= 10 {
		arr = append(arr, x%10)
		x /= 10
	}
	arr = append(arr, x)

	//compare head-tail
	l := len(arr)
	for i := 0; i < l+1/2; i++ {
		if arr[i] != arr[l-1-i] {
			return false
		}
	}
	return true
}

//复杂度为n
//计算倒置数，比大小。。。
func isPalindromeV1(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	origin := x
	y := 0

	for x >= 10 {
		y = 10*y + x%10
		x /= 10
	}
	y = 10*y + x

	return origin == y
}

//这次的思路是,还是使用取余倒置办法,x减 y增，用n表示变化次数,可能有2种情况
// 1. 位数为奇数,则首次y>x时, Xn == Y(n-1)
// 2. 位数为偶数,则y>x前, y==x
// 那么就好办了，变化过程中保留变化前的y >>> 「出现相等」或者「y>x且Xn == Y(n-1)」则true,否则false
// 复杂度n/2
// 【姑且先认为这个规模最小吧】，事实上在运行时间上没达到leetcode的top
func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 { //针对这种情况 21120
		return false
	}

	y0, y := 0, 0

	// zeroCount := 0
	for x >= 10 {
		y0 = y

		y = 10*y + x%10
		x /= 10

		if x == y {
			return true
		}

		if y > x {
			return y0 == x
		}
	}
	return false //事实上走不到这里
}

func main() {
	fmt.Println("hi")
	// isPalindromeV1(121)
	fmt.Println(21120%10 == 0)
}
