package main

import "fmt"

// 二、用户自定义都函数
type add func(a int, b int) int

// 三、高阶函数
// 高阶函数满足下列条件之一的函数：1、接收一个或多个函数作为参数。2、返回值是一个函数
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

// 在其他函数中返回函数
func simple1() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

// 闭包函数
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	// 一 、匿名函数
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T \n", a)

	// 二、用户自定义都函数类型
	var ad add = func(a int, b int) int {
		return a + b
	}
	ad(5, 6)
	fmt.Println("sum", ad(5, 6))

	// 三、高阶函数 把函数作为参数
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
	// 在函数中返回函数
	//s1 := simple1()
	//fmt.Println(s1(10,20))
	fmt.Println(simple1()(10, 20))

	// 闭包函数

	fmt.Println(appendStr()("world"))
	fmt.Println(appendStr()("Everyone"))
}
