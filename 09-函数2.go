package main

import "fmt"

// 函数调用流程：先调用后返回，先进后出
func funcb() (b int) {
	b = 100
	fmt.Println("funcb b = :", b)
	return
}
func funca() (a int) {
	a = 200
	b := funcb()
	fmt.Println("funca b = :", b)
	fmt.Println("funca a = :", a)
	return
}

// 递归函数 先调用本身 递归逻辑 先进后出 遇到条件才终止
func test(num int) {
	if num == 1 {
		fmt.Println("num=", num)
		return
	}
	test(num - 1)
	fmt.Println("num=", num)
}

// 通过递归实现数字累加 1+2+3...+100
// 方法一 for循环：
func test1() (num int) {
	for i := 0; i <= 100; i++ {
		num += i
	}
	return num
}

// 方法二 递归(升序)：
func test2(i int) (num int) {
	if i == 100 {
		return i
	}
	num = i + test2(i+1)
	return num
}

// 方法三 递归（降序）
func test3(i int) (num int) {
	if i == 1 {
		return i
	}
	num = i + test3(i-1)
	return num
}

// 回调函数 函数有一个参数是函数类型，这个函数就是回调函数
// 解决了面向对象的多态
type ABC func(int, int) int

func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

func Callback(a, b int, fTest ABC) (result int) {
	fmt.Println("回调函数...")
	result = fTest(a, b)
	return result
}

func main() {
	a := funca()
	fmt.Println("main a a = :", a)
	test(9)
	var sum int
	sum = test1()
	fmt.Println(sum)
	sum1 := test2(1)
	fmt.Println(sum1)
	sum2 := test3(100)
	fmt.Println(sum2)

	rest1 := Callback(222, 333, Minus)
	fmt.Println(rest1)
	rest2 := Callback(333, 222, Add)
	fmt.Println(rest2)

	// 调用带参数带匿名函数
	func(x, y int) {
		fmt.Println("x+y=", x+y)
	}(1000, 200)
}
