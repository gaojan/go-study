package main

import "fmt"

// 接口就是一个或多个方法签名的集合

// 接口定义：
type Humaner interface {
	// 方法，接口只有方法声明 没有实现，没有数据结构（由别的类型实现）
	sayhi()
}

type Student struct {
	name string
	id   int
	age  int
}

// Student类型实现了此方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d, %d] sayhi\n", tmp.name, tmp.id, tmp.age)
}

type Teacher struct {
	addr    string
	groupid string
}

// Teacher类型实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s,%s] sayhi\n", tmp.addr, tmp.groupid)
}

type Mystr string

// Mystr类型实现了此方法
func (tmp *Mystr) sayhi() {
	fmt.Printf("Mystr[%s] sayhi\n", *tmp)
}

func main() {
	// 定义接口类型的变量
	var i Humaner
	// 只是实现了此接口方法的类型，那么这个类型的变量（接受者类型）就可以给i赋值
	x := &Student{"mike", 666, 20}
	i = x
	i.sayhi()
	y := &Teacher{"shenzhen", "go"}
	i = y
	i.sayhi()
	var str Mystr = "hi"
	i = &str
	i.sayhi()
}
