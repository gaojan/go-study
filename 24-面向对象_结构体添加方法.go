package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

// 带有接收者的函数叫做方法
// 为结构体绑定方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp=", tmp)
}

// 通过一个函数给成员赋值  接收者接收一个指针
func (p *Person) SetInfo(n string, a int, s byte) {
	p.name = n
	p.age = a
	p.sex = s
}

func main() {
	p := Person{"xiaoming", 18, 'm'}
	p.PrintInfo()

	// 通过一个方法来修改结构体成员属性
	var p1 Person
	(&p1).SetInfo("yoyo", 10, 'm')
	p1.PrintInfo()
	p.PrintInfo()
}
