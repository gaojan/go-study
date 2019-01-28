package main

import "fmt"

// 位运算符 （一元运算符）
// >>  <<
//例子 运算符实现计算机储存单位的枚举
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("B:", B)
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
	fmt.Println("EB:", EB)
	fmt.Println("ZB:", ZB)
	fmt.Println("YB:", YB)

}
