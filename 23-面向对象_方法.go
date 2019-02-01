package main

import "fmt"

// 方法： 给函数绑定一个接收者的函数，叫做方法
// 可以给任意自定义类型添加方法（包括内置类型，不包括指针类型）
// 方法的定义：
// func (receiver receiverType) FuncName(param)(results)
// receiver 可以任意命名 如方法中未曾使用， 可省略参数名
// receiver 类型可以是T 或 *T（基类型T不能是接口或指针）
// 不支持重载方法。

// 面向过程
func Add(a, b int) int {
	return a + b
}

// 面向对象  方法：给某个类型绑定一个函数
type long int

// tmp叫接收者， 接收者就是传递的一个参数
func (tmp long) Add01(other long) long {
	return tmp + other
}

func main() {
	a := 10
	b := 20
	num := Add(a, b) // 普通函数调用方式
	fmt.Println("Add:", num)

	// 定义一个变量
	var x long = 20
	// 调用方法的格式： 变量名， 函数(所需参数)
	res := x.Add01(10)
	fmt.Println("res:", res)
}
