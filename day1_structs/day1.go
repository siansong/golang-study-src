package main

import (
	"fmt"
	"strings"
	"math"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

//pointers

func pointers() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

// struct
type Vertex struct {
	X int
	Y int
	ID string
}

func structFn() {
	v := Vertex{X: 233, Y: 666}
	v.X = 31415
	fmt.Println(v)

	// p := &v
	var p *Vertex = &v
	p.X = 233
	fmt.Println(v)

	// v1 := Vertex{}
}

// arrays
func arrayFn0() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	primes1 := [9]int{2, 3, 5, 7, 113}
	fmt.Println(primes1)

	s := primes[1:4]
	fmt.Println(s)
	fmt.Println(primes[:3])
	fmt.Println(primes[3:])
	fmt.Println(primes[:])

	s[0] = 33
	fmt.Println(s)
	fmt.Println(primes)
}

func arrFn1() {
	arr := [2]int{1, 2,}	
	fmt.Printf("%T, val: %v \n", arr, arr)
	
	q := []int{1, 2, 3, 4}
	fmt.Println(q)
	fmt.Printf("%T, val: %v \n", q, q)


	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(s)

	ss := s[:2];
	fmt.Println(len(ss), cap(ss))
	ss = s[:]
	fmt.Println(len(ss), cap(ss))

	var nilS []int
	if nilS == nil {
		fmt.Println("nil!!")
	}
}

func makeingSlices() {
	a := make([]int, 5)
	printSlice("a", a)

	// b := make([]int, 0, 5)
	b := []int{0, 1, 2, 3, 4}
	printSlice("b", b)

	c := b[1:2]
	printSlice("c", c)

	d := c[2:4]
	printSlice("d", d)//?? cap(d) 为什么不是5, counting from the first element in the slice.
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func sliceOfSlice() {
	board := [][]string{
		[]string{"-", "_", "-"},
		[]string{"(", "_", ")"},
	}
	fmt.Println(board)

	board[0][0] = "O"
	board[1][1] = "^"
	fmt.Println(board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], "$"))
		// fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceAppend() {
	var s []int
	printSlice("s", s)

	// append(s, 0)
	s = append(s, 0)
	printSlice("s", s)
	
	s = append(s, 1)
	printSlice("s", s)
	
	s = append(s, 2, 3, 4)
	printSlice("s", s)

}

func rangeFn0() {
	for i,v := range []int{1, 2, 4, 8} {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
	for i:= range pow {
		pow[i] = 1 << uint(i) 
	}

	for _, val := range pow {
		fmt.Printf("%d\n", val)
	}
}

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for i := range arr {
		arr[i] = make([]uint8, dx)
		// v = make([]uint8, dx)
		for j := range arr[i] {
			arr[i][j] = uint8((i + j)/2)
		}
	}
	return arr
}
func sliceExercise() {
	pic.Show(Pic)
}

func mapFn0(){
	m := make(map[string]int)
	m["foo"] = 233
	m["bar"] = 666.0
	fmt.Println(m)
	delete(m, "foo")
	foo, ok := m["foo"]
	fmt.Println(m, foo, ok)

	var m2 = map[int]int {
		123: 456,
		233: 666,
	}
	fmt.Println(m2)
}

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	var m map[string]int = make(map[string]int)
	for _, v := range arr {
		count := m[v]
		m[v] = count +1
	}
	return m
}

func mapsExecs(){
	wc.Test(WordCount)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func fp(){
	f := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(f(1, 2))
	fmt.Println(compute(f))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fpClosures() {
	pos, neg := adder(), adder()
	for i :=0; i < 10; i++ {
		fmt.Println(pos(i), neg(-i))
	}
}

func fibonacci() func() int {
	forward, backward := 0, 1 
	count := 0
	return func() int {
		count++
		if(count == 1) {
			return 0
		}
		tmp := forward + backward
		forward = backward
		backward = tmp
		return backward
	}
}

func fpFibonacci(){
	f := fibonacci()
	for i:=0; i<10; i++ {
		fmt.Println(f())
	}
}

func main() {
	// pointers()

	// structFn()

	// arrayFn0()

	// arrFn1()

	// makeingSlices()

	// sliceOfSlice()

	// sliceAppend()

	// rangeFn0()

	// sliceExercise()

	// mapFn0()

	// mapsExecs()

	// fp()

	// fpClosures()

	fpFibonacci()
}