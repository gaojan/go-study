package main

/*
一 无缓冲
无缓冲的channel 是指在接收前没有能力保存任何值的通道。
这种类型的channel要求发送goroutine和接收goroutine同时准备好，才能完成发送和接收操作。
如果两个goroutine没有同时准备好，通道会导致先执行发送或接收操作的goroutine阻塞等待。
这种对通道进行发送和接收对交互行为本身就是同步的。其中任意一个操作都无法离开另一个操作单独存在。

二 有缓冲
有缓冲的channel 是一种在被接收前能存储一个或者多个值的通道。
这种类型的channel并不强制要求goroutine之间必须要同时完成发送和接收。通道会阻塞发送和接收动作的条件也会不同。
只有在通道中没有接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞。
这导致有缓冲通道和无缓冲的通道之间的一个很大的不同：无缓冲的通道保证进行发送和接收的goroutine会这同一时间进行数据交换；
有缓冲的通道没有这种保证。

*/

import (
	"fmt"
	"time"
)

func main() {
	// 无缓冲 channel
	//var ch = make(chan int, 0)  // 0表示不能存储数据
	//fmt.Printf("无缓冲channel：len(ch)=%d, cap(ch)=%d\n", len(ch), cap(ch))

	// 有缓冲 channel
	var ch1 = make(chan int, 5) // 5表示可以存储5个元素
	fmt.Printf("有缓冲channel: len(ch1)=%d, cap(ch1)=%d\n", len(ch1), cap(ch1))

	//defer fmt.Println("主程序也结束")

	go func() {
		//defer fmt.Println("子协程调用完毕")
		for i := 0; i < 5; i++ {
			fmt.Println("子协程：i=", i)
			ch1 <- i
			fmt.Printf("子协程[%d]:len(ch)=%d, cap(ch)=%d\n", i, len(ch1), cap(ch1))
			//time.Sleep(time.Second)
		}
		//ch <- "我是子协程，工作完毕"
	}()
	time.Sleep(time.Second)
	//str := <- ch  //取数据 没有数据前阻塞
	//fmt.Println("str=", str)

	for i := 0; i < 5; i++ {
		num := <-ch1
		fmt.Println("主程序取数据：num=", num)
	}
}
