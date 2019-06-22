package main

import "fmt"

//字典

func main() {
	// 创建字典
	dit := map[int]string{1: "a"}
	fmt.Println("dict:", dit) //dict: map[1:a]
	mm := dit[1]
	fmt.Println("mm:", mm) // a
	// 添加元素
	dit[2] = "b"
	fmt.Println("dit:", dit) // dit: map[1:a 2:b]
	// 删除元素
	delete(dit, 1)
	fmt.Println("dit:", dit) //dit: map[2:b]

}
