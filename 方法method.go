package main

import "fmt"

// go虽然没有class 但依旧有method
type A struct {
	Name string
}
type B struct {
	Name string
}

// 定义结构体的方法
func (a *A) Print() { // a 局部变量 a变量接收结构体A
	a.Name = "AAAAA"
	fmt.Println("method A")
}
func (b B) Print() {
	b.Name = "BBBBB"
	fmt.Println("method B")
}

// 定义类型的方法 （struct是一类类型 int也是就类型 那么可以定义int类型的方法）
type TZ int

func (tz *TZ) Print() {
	fmt.Println("method TZ")
}

// 声明一个底层类型为int的类型，并实现调用某个方法就递增100
type C int

func (c *C) Increase(num int) {

	*c += C(num) // 将num转换成C类型
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)
	b := B{}
	b.Print()
	fmt.Println(b.Name)

	var tz TZ
	// 调用类型的方法一：
	tz.Print()
	// 调用类型的方法二：
	(*TZ).Print(&tz)
	// 方法可以访问私有属性和公有属性
	// 作业题目
	var c C
	c = 120
	c.Increase(100)
	fmt.Println(c)

}
