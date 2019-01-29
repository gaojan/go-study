package main

import "fmt"

// go函数不支持嵌套、重载和默认参数
// 支持以下特性：无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包

// 函数的定义
//func 函数名(参数1 类型， 参数2 类型)(返回值类型){ 没有返回值可以不写

//}
//不定长变参  ... 类型
func A(a ...int) {
	fmt.Println(a)
}

// ...int 类型一样的不定长参数， ...type不定长参数类型
// func Aa(args ...int){
// 		fmt.Println(len(args))  // len() 获取参数的个数
// }
func Aa(a string, b ...int) { // 如果多个不定长变参类型不一样 那么不定长参数一定写在最后面
	fmt.Println(a)
	fmt.Println(b)
}

// 多返回值函数
func B() (a, b, c int) {
	a, b, c = 8, 7, 6
	return a, b, c // 如果指定了返回值也可以不需要写返回 直接return
}

func Bb(s ...int) {
	s[0] = 3 // 传参的是a,b 但是得到的是值拷贝
	s[1] = 4
	//fmt.Printf("s:%v \n", s)  // s:[3 4]
}

func Bbb(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Printf("s:%v \n", s) //s:[5 6 7 8]
}
func Baa(a int) { //值类型不会改变 引用类型会发生改变， 如果要拷贝 进行指针传递
	a = 11
	fmt.Printf("a: %v \n", a)
}

func Bcc(b *int) { // 指针传递,可以改变值类型的值
	*b = 100
	fmt.Printf("b: %v, %v \n", b, *b)
}

// 函数也是一种类型 可以变量接收
func C() {
	fmt.Println("func C()")
}

// 函数类型 可以通过type给函数类型起名
type FuncType func(int int) int

// a函数参数传递给b函数
func mytest1(tmp ...int) {
	for _, data := range tmp {
		fmt.Printf("mytest1:%v \n", data)
	}
}
func mytest2(args ...int) {
	//mytest1(args ...)   // mytest2的值传递给mytest1
	mytest1(args[:2]...) // 可以利用切片只取前两个参数 包括终位置本身
}

// 匿名函数的使用
var N = func() { fmt.Println("这是个匿名函数") }

// 闭包函数 作用是返回一个函数
func Cc(x int) func(int) int { // 外部函数返回内部函数，内部函数返回一个int
	return func(y int) int {
		return x + y
	}
}

func main() {
	A(1, 2, 3, 4, 5, 6, 7, 8)
	Aa("hello", 9, 8, 7, 6)
	fmt.Println(B())
	a, b := 1, 2
	fmt.Printf("a, b: %v ,%v \n", a, b)
	Bb(a, b)
	s1 := []int{1, 2, 3, 4} // 引用类型，Bb 引用了s1的地址 所以s1地址改变 值改变 引用类型会发生改变 但是值类型不会改变
	Bbb(s1)
	fmt.Printf("s1:%v \n", s1) //

	a1 := 12
	Baa(a1)
	fmt.Printf("a1: %v \n", a1)

	b1 := 99
	Bcc(&b1)
	fmt.Printf("b1: %v \n", b1)
	fmt.Println(&b1)

	c := C // 函数类型的使用
	c()    // 调用c  func C()
	N()    // 匿名函数调用

	f := Cc(200)
	fmt.Println(f(100))
	fmt.Println(f(200))

	mytest2(100, 200, 300)
}
