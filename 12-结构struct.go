package main

import (
	"fmt"
)

// go 没有class 代替来class的位置 但没有class的功能 没有继承的概念
//struct定义： type <name> struct{}

// struct 定义
type person struct {
	Name string
	Age  int
}

// struct 嵌套 嵌套匿名struct
type pers struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

// struct的组合（类似继承，但没有继承的功能） 如：
type human struct { // teacher和student都human相同的属性
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

// 如果匿名字段和外层结构有同名字段，该如何操作？
type Test1 struct {
	Name string
}
type Test2 struct {
	Test1
	Name string
}

func main() {
	a := person{
		Name: "jack", // 跟下面 两种写法
		Age:  20,
	} // 有点类似类的实例化
	//a.Name = "jack"
	//a.Age = 20     // 跟类属性的操作是一样的
	fmt.Println(a)
	D(a)
	fmt.Println(a)

	aa := pers{
		Name: "hello world",
		Age:  30,
	}
	aa.Contact.Phone = "12345678"
	aa.Contact.City = "shenzhen"
	fmt.Println(aa)

	c := pers{} // c 也是pers类型
	c = aa
	fmt.Printf("c: %v \n", c)

	// 匿名结构 的定义
	b := &struct { // 加& 变成指针 取地址
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  23,
	}
	fmt.Println(b)

	t := teacher{Name: "老师", Age: 50, human: human{0}} // 组合写法一
	s := student{Name: "同学", Age: 10}
	//s.human.Sex = 1   // 组合写法二  这个写法跟下面的写法一样
	s.Sex = 1 // 跟上面写法类同
	fmt.Println(t)
	fmt.Println(s)

	tes := Test2{Name: "test2", Test1: Test1{Name: "test1"}}
	fmt.Println(tes)
	fmt.Println(tes.Name, tes.Test1.Name)
}

func D(per person) {
	per.Age = 13
	per.Name = "haha"
	fmt.Println("D", per)
}
