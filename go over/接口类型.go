package main

import "fmt"

//　接口 一个接口类型总是代表着某一种类型(既所有实现它的类型)
// 所谓实现一个接口中的方法是指，具有与该方法相同的声明并且添加了实现部分（由花括号包裹的若干条语句）
type Animal interface {
	Grow()
	Move(new string) (old string)
}

type Cat struct {
	name   string
	age    uint8
	status string
}

type Human struct {
	name   string
	Gender string
	Age    uint8
}

func main() {
	p := Human{"Robert", "Mala", 20}
	// 使用类型转换表达式把一个*Person类型转换成空接口类型的值：
	v := interface{}(&p)
	fmt.Println(v)
	myCat := Cat{"Little C", 2, "In the house"}

	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v \n", ok, animal)

}
