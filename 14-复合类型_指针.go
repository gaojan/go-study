package main

import "fmt"

/*
 分类：
类型       名称      长度     默认值    说明
pointer    指针              nil
array      数组              0
slice      切片              nil      引用类型
map        字典              nil      引用类型
struct     结构体            nil

指针特性：
1、默认值nil，没有NULL常量。
2、操作符"&"取变量地址，"*"通过指针访问目标对象
3、不支持指针运算，不支持->运算符，直接使用"."访问目标成员
*/

//方法一：
func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap: a=%d b=%d \n", a, b)
}

// 方法二：
func swap1(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("swap1: a=%d b=%d \n", *a, *b)
}

func main() {
	// 每个变量有2层含义：1、变量的内存。2、变量的地址
	var num int = 10
	fmt.Printf("a=%d \n", num)  // 变量的内存
	fmt.Printf("a=%v \n", &num) // 变量的地址 也叫指针
	// 保存某个变量的地址，需要指针类型  如：*int 保存int的地址，**int 保存 *int地址
	// 定义指针类型的变量：
	var P *int
	P = &num // 指针变量指向谁，就把谁的地址赋值给指针变量 &num表示把num的地址赋值给P
	fmt.Printf("P=%v \n &num=%v \n", P, &num)

	*P = 666                                 // *P 代表的是一个指针类型变量  操作的不是P的内存，而是P指向的内存（就是num）
	fmt.Printf("*p=%v\n num=%v \n", *P, num) // *p=666  num=666
	fmt.Printf("*p=%v\n num=%v \n", P, num)

	var O *int
	fmt.Println("O=", O)
	//*O = 9    // err O没有合法指向 报错  因为O为空指针 没有指向 需要赋值
	//fmt.Println(*O)
	// new()方法 可以动态分配内存地址 避免空指针
	O = new(int)
	*O = 9
	fmt.Println(*O)

	// 自动推导类型  自动识别q为指针类型
	q := new(int)
	*q = 111
	fmt.Println(q)
	fmt.Println(*q)

	// 变量交换
	// 方法一： 值传递
	a, b := 100, 200
	swap(a, b) // swap: a=200 b=100  内部的已经交换 a引用b的值 b引用a的值  外部的a b 不会改变
	// 方法二：地址传递
	swap1(&a, &b)                         // swap1: a=200 b=100  地址交换 外部的a b 也会改变
	fmt.Printf("main: a=%d b=%d\n", a, b) //main: a=200 b=100 外部的没有交换
}
