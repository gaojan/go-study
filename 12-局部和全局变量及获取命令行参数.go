package main

import (
	"fmt"
	"os"
)

var a int = 20 // 定义在函数外，全局变量  在任何地方都能使用

func A() {
	a = 20
	fmt.Println("外部a=", a) // a = 20
}

func main() {
	// os.Args 接收用户从命令行传来的参数，都是已字符串方式传递
	list := os.Args
	n := len(list)
	fmt.Println("n=", n)
	fmt.Println(list) // 第0个参数是程序本身， 第1个之后第参数是通过命令行传递到程序到参数

	// 局部变量
	{
		i := 111
		fmt.Println("i=", i) // 定义在{}里面的就是局部变量，只能在{}里面有效
		// 执行到定义变量那句话，才开始分配空间，离开作用域自动释放
	}
	//fmt.Println(i)  // 打印不来 会报错

	var a int // 局部变量  1、不同作用域，允许定义同名变量 2、使用变量的原则，就近原则
	a = 99
	fmt.Println("内部a=", a) // a=99
	A()
}
