package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("runoob.com")
	var a, b int = 10, 20
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)

}
