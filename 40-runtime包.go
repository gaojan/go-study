package main

import (
	"fmt"
	"runtime"
	"time"
)

// Gosched
//runtime.Gosched() 用于让出CPU时间片， 让出当前goroutine的执行权限，调度器安排其它等待的任务运行，并主下次某个时候从该位置恢复执行。

// Goexit
// runtime.Goexit() 将立即终止当前goroutine执行，调度器确保所有已注册defer延迟调用被执行

//GOMAXPROCS
// runtime.GOMAXPROCS() 用来设置可以并行计算的CPU核数的最大值，并返回之前的值

func test() {
	defer fmt.Println("cccccccccccc")
	//return  // 终止函数
	runtime.Goexit() // 终止所在的协程
	fmt.Println("dddddddddddddddd")
}

func Printer(tmp string) {
	for _, str := range tmp {
		fmt.Print(string(str))
		time.Sleep(time.Second)
	}
	fmt.Print("\n")
}

func person1() {
	Printer("hello")
}

func person2() {
	Printer("world")
}

func main() {
	// Gosched
	//go func() {
	//	for i := 0; i<5; i++{
	//		fmt.Println("子协程，i=", i)
	//	}
	//}()
	//
	//for i := 0; i<5; i++{
	//	// 让出时间片，先让别的协程执行，它执行完，再执行次协程
	//	runtime.Gosched()
	//	fmt.Println("main, i=", i)
	//}

	// Goexit
	//go func (){
	//	fmt.Println("aaaaaaaaaaa")
	//	test()
	//	fmt.Println("bbbbbbbbbbb")
	//}()

	// GOMAXPROCS
	//// 参数n:指定以单核运算
	//n := runtime.GOMAXPROCS(4)
	//fmt.Println("n=", n)
	//
	//for{
	//	go fmt.Print("1")
	//	fmt.Print("0")
	//}

	// 多任务资源竞争的问题
	// 新建两个协程，同时打印
	go person1()
	go person2()

	// 写个死循环，目的不让主程序结束
	for {

	}

}
