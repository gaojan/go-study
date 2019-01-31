package main

import "fmt"

// 数组做函数参数，是值传递，不会改变参数变量值，只在函数内部生效
// 实参数组的每个元素给行参数组拷贝一份 行参的数组是实参数组的值拷贝
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify: a=", a)
}

// b指向数组a，它是指向数组指针
// &a 代表指针所指向的内存，就是实参a
func modify1(b *[5]int) {
	b[0] = 666
	fmt.Println("modify1: b=", *b)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Println("main: a=", a)

	b := &a
	modify1(b)
	fmt.Println("main: ", a)
}
