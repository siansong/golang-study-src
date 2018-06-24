package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

//Note:
//类型声明 type {name} {type}
//结构体声明 type {name} struct
//函数声明 func {name}({params}) {returnType}
//	params> {name} {type}
//方法声明 func ({receiver}) {name}() {returnType}
//  receiver> {name} {type}
//接口声明
//type {name} interface {
// 	...methods signatures
// }
//接口实现,不需要显式的声明

func gDescribe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Vertex struct {
	X, Y float64
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func AbsFn(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// func AbsFn(v *Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
func (v Vertex) ValScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("valScale v:", v)
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("Scale v:", v)
}
func ScaleFn(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodFn0() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(AbsFn(v))
}

type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func methodFn1(){
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func methodFn2() {
	v := Vertex{3, 4}

	//部分语言里的值传递和引用传递
	v.ValScale(10)
	fmt.Println(v.Abs())
	
	v.Scale(10)
	fmt.Println(v.Abs())
}

func methodPointerExplain() {
	v := Vertex{3, 4}
	ScaleFn(&v, 10)
	// ScaleFn(v, 10)
	// AbsFn(&v)
	fmt.Println(AbsFn(v))
}


//Note: 使用指针作为接收者的2种情况，1.方法需要修改结构体；2.避免每次调用方法时都发生值拷贝（这点在接收者是一个大结构体时能更高效）


type Animal interface {
	Name() string
	MaxAge() int
}
type Human struct {
	name string
	maxAge int
}
func (h Human) Name() string{
	return h.name
}
func (h Human) MaxAge() int {
	return h.maxAge
}

func interfaceFn0() {
	var a Animal
	h := Human{"fooName", 100}
	a = h
	fmt.Println(a.Name())
}

type I interface {
	M() string
}
type Impl struct {
	X string
}
func (impl Impl) M() string {
	if &impl == nil {
		return "<nil>"
	}
	return impl.X
}
type Impl2 struct {}
func (i *Impl2) M() string {
	return "impl2.M"
}

func interfaceValuesWithNil() {
	var i I
	var impl *Impl
	i = impl
	fmt.Println(i.M())
}

func emptyInterface() {
	var i interface{}
	gDescribe(i)
	i = 42
	gDescribe(i)
	i = "233"
	gDescribe(i)
}

func typeAssert() {
	var i I
	impl := Impl{"x"}
	i = &impl
	fmt.Println(i.M())

	switch v := i.(type) {
	case *Impl:
		fmt.Println("type Impl")
	default:
		fmt.Println("default", v)
	}

	impl2 := Impl2{}
	i = &impl2
	// var i2 I = i.(*Impl2)
	i2, ok := i.(Impl)
	gDescribe(i2)
	fmt.Println(ok)
}

type IPAddr [4]byte
func (addr IPAddr) String() string {
	rs := fmt.Sprint(addr[0])
	for i := 1; i<len(addr); i++ {
		rs += "." + fmt.Sprint(addr[i])
	}
	return rs
}
func exerciseStringer() {
	// 测试类型转换
	fmt.Println(strconv.Itoa(666))
	hosts := map[string]IPAddr {
		"loopback": {127,0,0,1},
		"googleDNS": {8,8,8,8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// errors

type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func errFn0() {
	err := MyError{time.Now(), "it didn't work"}
	fmt.Println(err)
	fmt.Println(&err)

	i, _ := strconv.Atoi("233")
	// i, e := strconv.Atoi("233")
	fmt.Println(i)
}

func ioFn0() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b= %v \n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}
func (r MyReader) Read(b []byte) (i int, err error) {
	for i := range b {
		// it, _ := strconv.Atoi("A")
		// fmt.Println(it)
		b[i] = 65
	}
	return len(b), nil
}

func infiniteA() {
	reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (l int, e error) {
	// rot13map := map[string]string{
	// 	"a": "h",
	// }
	if &r == nil || r.r == nil {
		return 0, fmt.Errorf("<nil>")
	}
	br := make([]byte, len(b))
	l, e = r.r.Read(br)
	for i,_ := range br {
		// s, _ = rot13map[v]
		b[i] = 65
		//TODO rot13 cipher,这个算法懒得实现了。。。
	}
	return
}

func rot13ReaderFn() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func main() {
	// fmt.Println("Hello world")

	// methodFn0()
	// methodFn1()
	// methodFn2()
	// methodPointerExplain()

	// interfaceFn0()
	// interfaceValuesWithNil()
	// emptyInterface()
	// typeAssert()

	// exerciseStringer()

	// errFn0()

	// ioFn0()

	// infiniteA()

	rot13ReaderFn()
}