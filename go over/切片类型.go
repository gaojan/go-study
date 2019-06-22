package main

import "fmt"

// 切片
func main() {
	var number1 = [5]int{1, 2, 3, 4, 5}
	var slice1 = number1[1:]
	fmt.Printf("%v \n", slice1) // [2,3,4,5]
	var slice2 = slice1[1:2]
	fmt.Printf("%v \n", slice2)      //[3]
	fmt.Printf("%v \n", cap(slice2)) //3
	// 添加元素
	var slice3 = append(slice1, 6, 7)
	fmt.Printf("%v", slice3) // [2 3 4 5 6 7]
}
