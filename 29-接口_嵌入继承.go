package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Person interface {
	Humaner // 继承了Humaner接口的sayhi()
	sing(lrc string)
}
type Student struct {
	name string
	id   int
}

// Student 实现了一个sayhi方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s,%d]\n", tmp.name, tmp.id)
}
func (tmp *Student) sing(lrc string) {
	fmt.Printf("Stundent在唱歌：%s\n", lrc)
}

func main() {
	var p Person  // 超集（父类）
	var h Humaner // 子集（子类）
	//p = h  // err 子集不能转换为超集
	h = p // 超集可以转换为子集
	s := &Student{"jack", 9898}
	p = s
	p.sayhi() // Humaner继承过来的
	p.sing("hahaha")

}
