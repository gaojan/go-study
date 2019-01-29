package main

import "fmt"

// defer 在函数体执行结束后按照调用顺序的相反顺序逐个执行
// 即使函数发生错误也会执行  支持匿名函数的调用
// go中没有异常机制，但有panic/recover模式来处理错误

func main() {
	//for i:=0; i<5; i++ {
	//	defer fmt.Println(i)  // 4 3 2 1 0 相反顺序调用执行
	//}

	//for i:=0; i<5; i++ {
	//	defer func() {    // i作为参数传递进去，在匿名函数中实际上是引用来i的局部变量 所以在函数退出去的时候i的值为5 匿名函数指向5
	//		fmt.Println(i)    // 5 5 5 5 5
	//	}()  // 这个（）是调用这个匿名函数
	//}

	A1()
	B1()
	C1()

	//
	var fs = [4]func(){}
	fmt.Printf("fs: %T \n", fs)
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i =", i)                       //
		defer func() { fmt.Println("defer_closure i = ", i) }() //
		fs[i] = func() { fmt.Println("closure i = ", i) }       //在defer后面会先执行，得到四个匿名函数并存到fs的slice，传入的i值是通过外层函数传入，引用i的地址所以i=4
	}
	for _, f := range fs {
		f()
	} // 遍历fs 的fs中的f函数 调用
}

func A1() {
	fmt.Println("func A")
}

func B1() { // panic 异常函数
	defer func() { // 使用defer 捕获异常 会接着往下执行
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func C1() {
	fmt.Println("func C")
}
