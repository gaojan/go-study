package main

import "fmt"

// panic适用场景，当遇到不可恢复当错误状态当时候，如数组越界，空指针引用等，这些运行时错误会引起panic异常

func testa() {
	fmt.Println("aaaaaaaaaaa")
}

//func testb() {
//	//fmt.Println("bbbbbbbbbb")
//	// 显式调用panic函数，导致程序中断
//	panic("this is a panic test")
//}
// 数组越界导致panic
func testb(x int) {
	var a [10]int
	a[x] = 111 // 当x为20时候，导致数组越界，产生panic
	fmt.Println("x:", x, "a[x]:", a[x])
}

func testc() {
	fmt.Println("ccccccccccc")
}

func main() {
	testa()
	testb(10)
	testc()
}
