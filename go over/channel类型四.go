package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	defer close(ch)

}

func main() {
	// write 协程里会向ch写入0和1，接下来发生阻塞，直到ch内的值被读取。
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
	/*   容量为2 先进2个再出一个 先进先出
	successfully wrote 0 to ch
	successfully wrote 1 to ch
	read value 0 from ch
	successfully wrote 2 to ch
	read value 1 from ch
	successfully wrote 3 to ch
	read value 2 from ch
	successfully wrote 4 to ch
	read value 3 from ch
	read value 4 from ch
	*/
}
