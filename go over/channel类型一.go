package main

import (
	"fmt"
	"time"
)

// 通道channel类型
// 1、种类： 单向通道 双向通道 有缓冲通道 无缓冲通道
// 2、特点： 先进先出

func Hello(done chan bool) {
	fmt.Println("Hello world goroutine is going to sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("hello goroutine awake and going to write to done")
	done <- true
}

func main() {
	// 通道定义 用内建函数make 带缓存的通道
	ch1 := make(chan string, 5)
	ch1 <- "hello world"
	v1 := <-ch1
	fmt.Println(v1) // hello world
	// 无缓存的通道
	//ch2 := make(chan string, 0)
	//ch2 <- ""  // 不能放任何东西 回报错

	// 关闭通道
	close(ch1)

	// 单向通道--接收通道
	type Receiver <-chan string
	// 单向通道--发送通道
	type Sender <-chan string
	// 创建一个缓存通道 吧上面单向通道转化为双向通道
	myChannel := make(chan string, 3)
	myChannel <- "双向通道1"
	var sender Sender = myChannel
	v3 := <-sender
	fmt.Println(v3)
	myChannel <- "双向通道2"
	var receiver Receiver = myChannel
	v4 := <-receiver
	fmt.Println(v4)

	// 协程之间怎么通过通道通信
	done := make(chan bool)
	fmt.Println("main going to call hello go goroutine")
	go Hello(done)
	bl := <-done
	fmt.Println("main function")
	fmt.Println(bl)
}
