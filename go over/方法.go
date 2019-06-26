package main

import (
	"fmt"
	"math"
)

// 方法 方法其实就是一个函数， 在func这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者非结构体类型。
// 接收器是可以在方法的内部访问的

// 方法的定义
//f    接收器类型    方法名
//func (t Type) methodName(param Type)  { }

type Employee struct {
	name string
	age  int
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

// 属于结构体的匿名字段都方法可以被直接调用，就好像这些方法是属于定义来匿名字段的结构体一样
type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Println(a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

// displayAge()方法将Employee作为接收器类型  值接收器
func (e Employee) displayAge() {
	fmt.Println("名字：", e.name, "年龄：", e.age)
}

//displayAge1()方法被转化为一个函数，吧Employee当中参数传入
func displayAge1(e Employee) {
	fmt.Println("名字：", e.name, "年龄：", e.age)
}

// 给Rectangle添加方法 Area 求面积方法
func (r Rectangle) Area() int {
	return r.length * r.width
}

// 给Circle 添加方法 Area 求圆周率方法
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// 使用值接收器
func (e Employee) changName(newName string) {
	e.name = newName
}

// 使用指针接收器
func (e *Employee) changAge(newAge int) {
	e.age = newAge
}

func main() {
	emp1 := Employee{"jack", 20}
	emp1.displayAge() // 调用Employee类型的displayAge()方法
	displayAge1(emp1)
	// displayAge方法被转化为一个函数，Employee结构体被当作参数传递给它。
	// 既然可以使用函数写出相同都效果，那为什么我们需要方法？有如下几点原因
	// 1、Go不是纯粹都面向对象编程语言，而且Go不支持类。因此基于类型的方法
	// 是一种实现和类相似行为的途径。
	// 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。
	// 假设我们有一个Square和Circle结构体。可以在Square和Circle上分别定
	// 义一个Area方法:
	r := Rectangle{10, 5}
	fmt.Println("长方形的面积是：", r.Area())

	c := Circle{12}
	fmt.Println("圆周率为:", c.Area())

	// 指针接收器和值接收器
	e := Employee{"Mark", 20}
	fmt.Println("Employee name 值接收器改变前：", e.name) // Mark
	fmt.Println("Employee age 值接收器改变前：", e.age)   // 20
	e.changName("Like")
	fmt.Println("Employee name 值接收器改变后：", e.name) // Mark
	e.changAge(50)                                // e.changAge(50) 可以代替（&e）.changAge(50)
	fmt.Println("Employee age 指针接收器改变后：", e.age)  // 50

	// 匿名字段的方法
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			"shenzhen",
			"guangzhou",
		},
	}
	p.fullAddress()

}
