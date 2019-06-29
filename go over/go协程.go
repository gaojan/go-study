package main

import (
	"fmt"
	"time"
)

/*
Go 编程原生支持并发。Go使用协程（Goroutine）和信道（Channel）来处理并发。
Go协程是什么？
Go协程是与其他函数或方法一起并发运行都函数或方法，Go协程可以看作是轻量级线程，创建一个Go协程都成本很小。
Go协程相比于线程的优势：
1、Go协程成本极低，堆栈大小只有若干kb。
2、Go协程会复用数量更少的os线程
3、Go协程使用信道（channel）来进行通信。信道用于防止多个协程访问共享内存时发生竞态条件，信道可以看作是Go协程之间通信的管道。
*/

func sayHello() {
	fmt.Println("hello world!")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d \n", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c \n", i)
	}
}

func main() {
	// 启动一个go协程
	go sayHello() // 启动一个新协程时，协程的调用会立即返回，与函数不同，程序控制不会去等待go协程执行完毕，在调用go协程之后，
	// 程序控制会立即返回到代码到下一行，忽略该协程到任何返回值
	// 如果希望运行其他Go协程，主协程必须继续运行着，如果主协程终止，则程序终止，于是其他协程也不会继续运行。
	//程序控制没有等待 hello 协程结束，立即返回到了代码下一行，打印 main function。
	// 接着由于没有其他可执行的代码，Go 主协程终止，于是 hello 协程就没有机会运行了。
	// 修复问题：
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	// 启动多个go协程
	go numbers()
	go alphabets() //1 a 2 3 b 4 c 5 d e main terminated
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
