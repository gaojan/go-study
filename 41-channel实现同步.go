package main

import (
	"fmt"
	"time"
)

/*
goroutine奉行通过通信来共享内存，而不是共享内存来通信
channel类型 和map类似名，channel也一个对应make创建的底层数据结构的引用。
channel 可以看作是管道或通道
默认情况下 channel接收和发送数据都是阻塞的。
当 capacity=0 时，channel是无缓冲阻塞读写的，当capacity>0 时，channel有缓冲，是非阻塞当，直到写满capocity个元素才阻塞写入。

*/

// channel通道创建
// channel顺序：取数据<- ch  写数据 ch <-
var ch = make(chan int) //ch都类型是 chan int

func Printer(tmp string) {
	for _, str := range tmp {
		fmt.Print(string(str))
		time.Sleep(time.Second)
	}
	fmt.Print("\n")
}

func person1() {
	Printer("hello")
	ch <- 8989 // 给管道写数据(随便放什么数据)， 发送数据
}

func person2() {
	<-ch // 从管道取数据，接收，如果通道没有数据它就会阻塞
	Printer("world")
}

func main() {
	// 多任务资源竞争的问题
	// 新建两个协程，同时打印
	go person1()
	go person2()

	// 写个死循环，目的不让主程序结束
	for {

	}

}
