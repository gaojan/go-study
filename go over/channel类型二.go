package main

import "fmt"

/*
计算一个数中每一位都平方和与立方和，然后把平方和与立方和相加并打印出来
*/

// 计算平方
func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10 // /= 相除后再赋值，C /= A 等于 C = C / A
	}
	squareop <- sum
}

// 计算立方
func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

// 遍历信道 并关闭信道
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println(squares + cubes)

	// 死锁
	// 因为没有其他协程从ch接收数据，触发panic
	//ch := make(chan int)
	//ch <- 5

	// 关闭信道和使用for range遍历信道
	// 数据发送发可以关闭信道，通知接收方这个信道不再有数据发送过来
	ch := make(chan int)
	go producer(ch)
	//for {
	//	v, ok := <- ch
	//	if ok == false{  // ok=false说明信道已经关闭
	//		break
	//	}
	//	fmt.Println("Received", v, ok)
	//}
	// 用for range重写上面方法
	for v := range ch {
		fmt.Println("Received", v) // for range 循环从信道 ch 接收数据，直到该信道关闭。一旦关闭了 ch，循环会自动结束。该程序会输出：
	}
}
