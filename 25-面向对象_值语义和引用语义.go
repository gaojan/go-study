package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

// 接收者为普通变量， 非指针，值语义， 相当于一份拷贝
func (p Person) SetInfoValue(x string, y int, z byte) {
	p.name = x
	p.age = y
	p.sex = z
	fmt.Println("SetInfoValue: &p=", &p)
}

// 接受者为指针变量， 引用传递
func (p *Person) SetInfoPointer(x string, y int, z byte) {
	p.name = x
	p.age = y
	p.sex = z
	fmt.Println("SetInfoPointer: p=", p)
}

func main() {
	s1 := Person{"xiaohua", 20, 'm'}
	fmt.Println("main: &s1=", &s1)
	// 值语义
	s1.SetInfoValue("mike", 10, 'm')
	fmt.Println("值传递后: &s1=", &s1) // 方法内部形参改变 外部实参未改变
	// 引用语义
	s1.SetInfoPointer("joe", 100, 'f') // 方法内部行参改变， 外部实参也跟着改变
	fmt.Println("引用传递后：s1=", s1)
}
