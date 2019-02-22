package main

import "fmt"

// recover必须在defer调用的函数中有效
// recover不会导致程序崩溃， 异常捕获

func testa() {
	fmt.Println("aaaaaaa")
}
func testb(x int) {
	defer func() { // recover()可以打印panic的错误信息
		if err := recover(); err != nil { // 产生panic异常
			fmt.Println(err)
		}
	}()
	var s [10]int
	s[x] = 111
	fmt.Println(s[x])
}
func testc() {
	fmt.Println("bbbbbbbb")
}

func main() {
	testa()
	testb(10)
	testc()
}
