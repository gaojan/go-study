package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
	addr string
}

// 值传递
func test1(s student) {
	s.id = 666
	s.name = "gaoxing"
	s.age = 12
	s.addr = "shanghai"
	fmt.Println("test1中：s=", s)
}

// 指针传递
func test2(p *student) {
	p.id = 666
	p.name = "gaoxing"
	p.age = 12
	p.addr = "shanghai"
	fmt.Println("test2中：p=", p)
}

func main() {
	s := student{1, "gao", 20, "shenzhen"}
	fmt.Println("入参前：s=", s)
	// 值传递， 形参无法改实参 实际上是拷贝了一份 形参修改拷贝的那份
	test1(s)
	fmt.Println("入参后：s=", s)
	// 指针传递 (引用传递)，指针指向实参的地址，形参改变，实参也会跟着改变
	p := &student{1, "gao", 20, "shenzhen"}
	fmt.Println("入参前：p=", p)
	test2(p)
	fmt.Println("入参前：p=", p)

}
