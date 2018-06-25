package main

import (
	"fmt"
	"time"
)

func fn0() {
	go fmt.Println("world!")
	fmt.Println("Hello ")
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
	}
}

func sum(s []int, c chan int) {
	fmt.Println(s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func chanFn0() {
	s := []int{1, 2, 3, 4}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	fmt.Println("go")
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func bufferedChan() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	count := 0
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Printf("quit, and default count: %d \n", count)
			return
		default:
			count ++
		}
	}
}

func selectFn() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 13; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

// RunGoroutines ... 文件内"main"
func RunGoroutines() {
	// fn0()
	// chanFn0()
	// bufferedChan()
	// rangeAndClose()
	selectFn()
}
