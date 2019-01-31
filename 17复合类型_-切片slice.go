package main

import "fmt"

// slice ：= make([]type, len, cap)
/*
切片: [low:high:max]
low: 下标的起点
high: 下标的终点(不包括下标本身)，[a[low], a[high]]左闭右开
长度： len = high - low
容量： cap = max - low
*/

// 切片的长度和容量
func test() {
	arr := []int{1, 2, 3, 0, 0} // 切片初始化
	s := arr[0:3:5]             // [low:high:max]
	fmt.Println("s=", s)
	fmt.Println("len=", len(s)) // 长度 3-0
	fmt.Println("cap=", cap(s)) // 容量 5-0

	s1 := arr[1:4:5] // [low:high:max] 从1开始切片
	fmt.Println("s1=", s1)
	fmt.Println("len1=", len(s1)) // 长度 4-1
	fmt.Println("cap1=", cap(s1)) // 容量 5-1
}

func main() {
	// 切片本身并不是数组或指针，它通过内部指针和相关属性引用数组片段，它指向底层的数组
	// slice和array数组的区别
	var ar [5]int // 数组[]中是固定常量数字，数组不能修改长度和容量
	var sl []int  // 切片[]里面为空，或者为...  切片长度和容量可以不固定，如不指定容量，容量和长度一样
	fmt.Println(ar)
	fmt.Println(sl)
	// slice 获取数组元素
	a := [9]int{8: 8}
	s1, s2 := a[8], a[6] // 一个slice只取一个元素
	fmt.Println(a)
	fmt.Println(s1, s2)
	b := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	sa := b[1:5] // 切片包含其位置 不包含始位置
	sb := sa[0:5]
	fmt.Println("sa", sa)
	fmt.Println(sb)

	// append函数会智能的底层数组的容量增长，一旦超过原底层数组容量，通常以2倍容量重新分配底层数组，并复制原来的数据

	n1 := make([]int, 3, 6)
	n2 := append(n1, 2, 3)
	fmt.Printf("%v, %p \n", n2, n2) // n2和n3内存地址一样 因为n2,n3未超出slice n1的容量 所以返回原始slice
	n3 := append(n1, 1, 2, 3)
	fmt.Printf("%v, %p \n", n3, n3)

	n4 := append(n1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // n4跟n2 n3不同地址，因为超过slice的容量，已重新分配数组并拷贝原始数据
	fmt.Printf("%v, %p \n", n4, n4)

	n5 := n1[2:5]
	n6 := n1[1:3]
	n5[0] = 9           // n5 和n6为同一个内存地址 改变n5或n6中的某个元素，会跟着一起改变
	fmt.Println(n5, n6) // [9 1 2] [0 9]

	//copy函数
	//copy(目标元素， 被拷贝元素)
	m := make([]int, 4, 5)
	//m := []int{0,0,0,0}
	fmt.Printf("%T,%v \n", m, m)
	m1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%T,%v \n", m1, m1)

	m2 := []int{8, 9, 0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("%T,%v \n", m2, m2)

	// copy(目的，源)
	copy(m1, m2) //copy会覆盖原来的元素，大容量slice copy到小slice 会以小slice容量为准
	fmt.Println(m1)

	// 此方法也可以完成
	l := []int{1, 2, 3, 4, 5}
	//l1 := l  // 赋值copy不改变地址
	//l1 := l[0:5]
	l1 := l[:]
	fmt.Printf("%p, %v \n", l, l)
	fmt.Printf("%p, %v \n", l1, l1)
	test()

}
