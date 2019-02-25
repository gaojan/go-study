package main

import (
	"fmt"
	"time"
)

/*
并行(parallel): 指在同一时刻，有多条指令在多个处理器上同时执行。
并行上两个队列同时使用两台咖啡机

并发(concurrency)：指在同一刻只能有一条指令执行，但多个进程指令被快速的轮换执行，
使得宏观上具有多个进程同时执行的效果，但在微观上并不是同时执行的，只是把时间分成若干段，使多个进程快速交替的执行。
并发是两个队列交替使用一台咖啡机

Go从语言层面就支持了并发，因为Go语言提供了自动垃圾回收机制

goroutine
goroutine 是Go并行设计的核心，goroutine其实就是协程，但是它比线程更小，执行goroutine只需极少的栈内存（4-5kb），
当然会根据相应的数据伸缩。正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

*/

func newTask() {
	for {
		fmt.Println("This is new task!!!!")
		time.Sleep(time.Second) // 延时1秒
	}
}

func main() {
	// goroutine 创建协程 使用go关键字
	go newTask()
	for {
		fmt.Println("This is main task!!!!")
		time.Sleep(time.Second) // 延时1秒
	}
}
