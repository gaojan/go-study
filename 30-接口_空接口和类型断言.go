package main

import (
	"fmt"
)

// 空接口不包含任何的方法，所有所有的类型都实现了空接口，因此空接口可以存储任意类型都数值。
// ... 可变参数 空接口类型，可以接收任意类型的参数0到多个
func Arg(arg ...interface{}) {
	fmt.Print("arg:", arg)
}

type Student struct {
	name string
	id   int
}

func main() {
	// 万能接口类型，保存任意类型都值
	var i interface{} = 1
	var s interface{} = "hello"
	fmt.Printf("i:%d\n", i)
	fmt.Printf("s:%s\n", s)
	i = "abc"
	fmt.Printf("i:%s\n", i)

	x := make([]interface{}, 3) // 定义一个切片
	x[0] = 1
	x[1] = "hello world"
	x[2] = Student{"jack", 666}
	for index, data := range x {
		// 第一个返回的是遍历的值，第二个返回判断结果的真假
		// 通过if实现断言
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d]类型为int,值为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d]类型为string,值为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d]类型为Student,值为：name=%s,id=%d\n", index, value.name, value.id)
		}
	}
	fmt.Println("-----------------------")
	// 通过switch实现断言
	for index, data := range x {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]类型为int,值为%d\n", index, value)
		case string:
			fmt.Printf("x[%d]类型为string,值为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d]类型为Student,值为：name=%s,id=%d\n", index, value.name, value.id)
		}
	}
}
