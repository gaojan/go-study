package main

import "fmt"

//
//const (
//	TOP int = 10
//)

// 操作符“&”取变量地址，“*” 通过指针间接访问目标对象
//var a=6
//var b *int=&a

func main() {
	//fmt.Println(&a) //0x5461e8
	//fmt.Println(*b)  //6
	//fmt.Println()
	//a :=1
	//for {
	//	a++  // 递增递减  ++  -- 是作为语句而不是作为表达式 所以只能单独一行
	//	fmt.Printf("%d--", a)
	//}

	//九九乘法表
	//for i:=1; i<TOP; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d * %d=%d  ", i, j, i*j)
	//	}
	//	fmt.Println()
	//}

	// if语句中 局部变量的作用域
	a := 10
	if a := 2; a > 1 {
		fmt.Println(a) // 2
	}
	fmt.Println(a) // 10
}
