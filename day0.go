package main

import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
	"runtime"
	"time"
)

/** 用函数来代替教程里的文件名了 */


/** basics start **/
func hello() {
	fmt.Println("hello, 世界")
}

func packages() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}

func exportedNames() {
	fmt.Println(math.Pi);
}

func add (x ,y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return 
}

//variables
var c, python, java bool
func variables() {
	var i int
	fmt.Println(i, c, python, java)
}

// variables with initializers
var i, j int = 233, 666
func varsWithInitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func shortVarsDeclare() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

// basic types
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128  = cmplx.Sqrt(-5 + 12i)
)
func basicTypes() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// struct ZeroValStruct{//233 已经忘完了
// 	i int
// }
func zeroVal() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//TODO struct nil??
}

func typeConvert() {
	var x, y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// var v uint = f
	// fmt.Println(v)
}

func typeInference() {//类型推导
	v := 42
	fmt.Printf("v is of type %T \n", v)

	y := 3.14
	fmt.Printf("y is of type %T \n", y)

	z := 0.5i
	fmt.Printf("z is of type %T \n", z)
}


const MyPi = 3.1415926
func constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", MyPi, "Day")

	// World = "world"

	const Truth = true
	fmt.Println("Go rules?", Truth)
}


// numberic constants
const (
	Big = 1 << 100
	Small = Big >> 99
)
func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func numConst() {
	// fmt.Println("Big : %v, Small %v", Big, Small)
	fmt.Println(Small)
	// fmt.Println(1<<100)
	fmt.Println(1<<1)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/** basics end */


/**  flowcontrol start */
func forFlow() {
	sum := 0
	for i :=0; i< 10; i++ {
		sum  += i
	}
	fmt.Println(sum)


	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println(sum2)
}

func gosWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func forever() {
	for {
		fmt.Println("233 666")
	}
}


func ifSyntax(){
	if 2 < 1{
		fmt.Println("2> 1")
	} else {
		fmt.Println("else")
	}
	
	if v := math.Pow(2, 3); v < 9 {
		fmt.Println(" 8 < 9")
	}
}


func mySqrt1(x float64) float64 {//ver1 暴力技术，精确到整数
	if x<0 {
		fmt.Println("sqrt x should be larger than 0")
		return x
	}

	i := 1.0
	for ;(i+1)*(i+1) < x; i++ {
	}
	fmt.Println("sqr1:根号", x , "~=", i)
	return i
}

func myAbs(x float64) float64 {
	if x<0 {
		return -1 * x
	}else {
		return x
	}
}

func mySqrt2(x float64) float64 {//ver2 按照提示写
	/**
	这个公式,我们假定y是我们要求的值
	z -= (z*z - x) / (2*z)
	我们假设z是随意的正浮点数，如果z>y,则z经过-=后值减，否则值增。每次的变化值可以大致理解为近似度。。。。
	*/
	z := 1.0
	tmp := 0.0
	count :=0
	for ; myAbs(tmp - z) > 0.000000001; z -= (z*z - x) / (2*z) {
		tmp = z
		count++
		fmt.Println("mySqrt2 internal z:", z)
	}
	fmt.Println("count: ", count, tmp)
	return z
}

func loopExercise() {
	mySqrt1(-1.0)
	mySqrt1(5)//3
	fmt.Println(math.Sqrt(5))

	fmt.Println("mySqrt2  ", mySqrt2(15))
	fmt.Println("math.Sqrt", math.Sqrt(15))
}

func switchSynctax() {
	fmt.Println(runtime.GOOS)
	switch os := runtime.GOOS; os {
	case "notPossible": 
		fmt.Println("Not psb")
	case "darwin":
		fmt.Println("OS X.")
		fallthrough //去到下个case
	case "linux":
		fmt.Println("linux.")
		// fallthrough
	default:
		fmt.Printf("%s.", os)
	}
}

func switchEvalOrder() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func testDeffer2() {
	defer fmt.Println("world 2")

	fmt.Println("Hello(2)")
}
func testDefer() {
	defer fmt.Println("world")

	defer testDeffer2()
	// testDeffer2()

	fmt.Println("Hello")
}
/**  flowcontrol end */

func main() {
	// packages()	

	// exportedNames()
	
	// fmt.Println(add(233, 667))

	// a,b := swap("hello", "world")
	// fmt.Println(a, b)

	// fmt.Println(split(17))
	// x,y := split(18)
	// fmt.Println(x, y)


	// variables()

	// varsWithInitializers()

	// shortVarsDeclare()

	// basicTypes()

	// zeroVal()

	// typeConvert()

	// typeInference()

	// constants()

	// numConst()


	//>>>flow control start
	// forFlow()

	// gosWhile()

	// forever()

	// ifSyntax()

	// loopExercise()

	// switchSynctax()

	// switchEvalOrder()

	// switchWithNoCondition()

	testDefer()
}