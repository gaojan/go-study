package main

import (
	"fmt"
	"time"
)

// 主协程先退出，其它子协程也要跟着退出
// 主协程先退出导致子协程没来的及调用

func main() {
	// 子协程
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i=", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main i=", i)
		time.Sleep(time.Second)
		if i == 5 {
			break
		}
	}

}
