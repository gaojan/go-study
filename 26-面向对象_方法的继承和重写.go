package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p *Person) PrintInfo() {
	fmt.Printf("name = %v \n sex = %v \n age = %v \n", p.name, p.sex, p.age)
}

// 继承Person的所有成员和方法
type Student struct {
	Person // 匿名字段
	id     int
	addr   string
}

// Student也实现了一个方法，这个方法和Person方法同名，次方法也叫重写
func (p *Student) PrintInfo() {
	fmt.Println("student: p=", p)
}

func main() {
	p := Person{"hello", 'm', 12}
	p.PrintInfo()
	s := Student{p, 11, "sz"}
	//fmt.Println(s)
	s.PrintInfo() // 方法重写 就近原则， 先找本作用域的方法，找不到再找继承的方法
	// 显示调用继承的方法
	s.Person.PrintInfo()
}
