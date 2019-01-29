// 程序所属包
package main

// 导入包
import "fmt"

// 单行注释
/*
多行注释
多行注释
*/

// 常量定义，一般大写 string 可写可不写 最好注明
const NAME string = "imooc"

// 变量定义  大驼峰
var Name string = "helloworld"

// 同时声明多个变量 并初始化
var a, b, c = "1", "2", "3"

// 自动推到型 必须初始化，通过初始化的值确定类型
//d := 2
// 自动推动类型和赋值的区别
// 赋值， 赋值前 必须先声明变量 var a string  赋值可以使用n次
// := 自动推动类型，先声明类型，再赋值， 同一个变量名只能使用一次，用于初始化那次

func abc() {
	println(a, b, c)
}

// 特殊只写变量“_”，用于忽略值占位 也叫匿名变量
func test() (int, string) {
	return 1, "abc"
}

// 一般类型声明
type imoocInt int

// 结构的声明
type Learn struct {
}

// 声明接口
type Ilearn interface {
}

// 函数的定义
func haha() {
	fmt.Println(111)
}

func main() {
	haha()
	fmt.Println("hello world!")
	fmt.Print(NAME)
	fmt.Print(Name)
	abc() // 123
	_, s := test()
	println(s)
	// %T 打印格式化类型
	fmt.Printf("a:%T", a)

	// 交换2个变量
	n, m := 1, 2
	var tmp = n
	n = m
	m = tmp
	fmt.Println(n)
	fmt.Println(m)
}
