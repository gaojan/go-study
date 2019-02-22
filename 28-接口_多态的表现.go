package main

import "fmt"

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

func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	s := &Student{"jack", 333, 20}
	t := &Teacher{"shenzhen", "python"}
	var str Mystr = "hello world"
	// 调用同一个函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	// 创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, i := range x {
		i.sayhi()
	}
}
