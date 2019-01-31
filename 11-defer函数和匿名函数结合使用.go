package main

import "fmt"

func main() {
	a := 10
	b := 20
	//defer func(){
	//	fmt.Println("内部：a=", a, "b=\n", b)  // 后调用，使用a、b的值改变，而不是a、b变量初始值
	//}()   // a=11  b=12
	defer func(a, b int) {
		fmt.Println("内部：a=", a, "b=\n", b) // 后调用，使用a、b的值改变，而不是a、b变量初始值
	}(a, b) // 把参数传递过去，已经先传递了参数，只是没有调用  //  a=10 b=20
	a = 11
	b = 12
	fmt.Println("外部：a=", a, "b=\n", b)
}
