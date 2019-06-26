package main

import "fmt"

// 结构体可以封装属性和方法

type Person struct {
	name   string
	Gender string
	Age    uint8
}

// 为Person结构图添加了一个方法
func (person *Person) Grow() {
	person.Age++
}

func (person *Person) Apd() {
	person.name = person.name + "_a"
}

func main() {
	p := Person{"Robert", "Male", 20}
	p.Grow()
	p.Apd()
	fmt.Println(p)

	// 匿名结构体积
	p1 := struct {
		Name string
		Age  uint8
	}{"gao", 20}
	fmt.Printf("%v", p1)

}
