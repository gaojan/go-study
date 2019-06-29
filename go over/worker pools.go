package main

import (
	"fmt"
	"sync"
	"time"
)

// 缓冲信道的重要应用之一就是实现工作池
// 一般而言，工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。
// WaitGroup用于实现工作池, 用于等待一批Go协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended \n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
