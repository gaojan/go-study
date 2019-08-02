package main

import (
	"fmt"
)

//type Mystruct struct {
//	Name string
//	First string
//	Age  int
//}

type Mystruct struct {
	str string
}

type CountStr interface {
	Count() int
}

func (s Mystruct) Count() int {
	st := s.str
	//if st == "" {
	//	return 0
	//}else {
	//	var num int
	//	for i := 0; i < len(st); i++ {
	//		fmt.Println("字符串的长度为：", i)
	//		num = i+1
	//	}
	return len(st)
}

//func printer(c []CountStr) {
//	fmt.Println("字符串的长度为：", c)
//}

func main() {
	//n := Mystruct{}
	//immutable := reflect.ValueOf(n)
	//val1 := immutable.FieldByName("First").String()
	//if val1 == "" {
	//	fmt.Println("有值")
	//	fmt.Printf("值是:%s \n", val1)
	//}else {
	//	fmt.Println("没有值")
	//}
	my := Mystruct{"gaojian"}
	//my1 := Mystruct{"helloworld"}
	c := []CountStr{my}
	//printer(c)
	fmt.Println(c[:])
}
