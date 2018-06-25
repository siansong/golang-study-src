package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		fmt.Println("nil tree")
	}
	//left
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	//middle
	ch <- t.Value
	//right
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := 0; i < 10; i++ {
		v1, v2 := <-c1, <-c2
		if v1 != v2 {
			return false
		}
	}
	return true
}

// BinaryTreeMain main
func BinaryTreeMain() {
	t := tree.New(1)
	t1 := tree.New(1)
	t2 := tree.New(2)

	c1 := make(chan int, 10)
	Walk(t, c1)
	// for i := 0; i< 10; i++{
	// 	fmt.Println(<- c1)
	// }

	fmt.Println(Same(t, t1))//true
	fmt.Println(Same(t, t2))//false
}