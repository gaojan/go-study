package main

import (
	"fmt"
	"time"
)

// 如果存在多个channel的时候，我们该如何操作呢，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。
// select 语句用于在多个发送/接收信道操作中进行选择。select语句会一直阻塞，直到发送/接收操作准备就绪。
// select语法与switch类似,所不同的是，这里的每个case语句都是信道操作。

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value:", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}

// 空select
//func main()  {
//	select{}
//}
// 除非有case执行，select语句就会一直阻塞着，如果select语句没有任何case，因此他会一直阻塞，导致死锁，触发panic
