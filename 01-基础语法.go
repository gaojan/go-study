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

// 多变量赋值
var a, b, c = "1", "2", "3"

func abc() {
	println(a, b, c)
}

// 特殊只写变量“_”，用于忽略值占位
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

// Println 换行打印
// Print 不换行打印

// 函数的定义
func haha() {
	fmt.Println(111)
}

// 一个go文件中 必须要有main 函数，go文件的入口，否则不能执行
func main() {
	haha()
	fmt.Println("hello world!")
	fmt.Print(NAME)
	fmt.Print(Name)
	abc() // 123
	_, s := test()
	println(s)
}
