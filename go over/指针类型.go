package main

import "fmt"

// 指针是一种存储变量内存地址的变量
// &操作符用于取变量地址

// 向函数传递指针参数
func change(val *int) {
	*val = 55
}

// 不要向函数传递数组的指针，而应该使用切片
func modify(arr *[3]int) {
	//(*arr)[0] = 90
	// a[x] 是(*a)[x]的简写形式， 因此上面代码中的(*arr)[0]可以替换为arr[0]
	arr[0] = 90
}

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("a的类型是：%T \n", a)
	fmt.Println("b的地址是：", a)

	// 一、指针的零值
	aa := 25
	var bb *int
	if bb == nil {
		fmt.Println("bb:", bb) // 任何未赋值前的指针的零值是nil
		bb = &aa
		fmt.Println("bb after:", bb) // 0xc00001a078
	}

	// 二、指针的解引用：将a解引用的语法是*a
	b1 := 255
	a1 := &b1
	fmt.Println("b1的地址是：", a1) // 0xc00001a090
	fmt.Println("b1的值是：", *a1) // 255
	// 用指针修改b1值
	*a1++
	fmt.Println("b1修改后的值为：", b1) // 256

	// 三、向函数传递指针参数
	a2 := 58
	fmt.Println("函数调用前a2的值是：", a2) // 58
	b2 := &a2
	change(b2)
	fmt.Println("函数调用后b2的值是：", a2) // 55

	// 四、不要向函数传递数组的指针，而应该使用切片
	a3 := [3]int{89, 90, 91}
	modify(&a3)
	fmt.Println("a3:", a3) //[90 90 91]  修改了 arr[0]的值
	// a[x] 是(*a)[x]的简写形式， 因此上面代码中的(*arr)[0]可以替换为arr[0]

	// 五、Go不支持指针运算
	//b3 := [...]int{100,200,201}
	//p :=&b3
	//p ++        // 会报错，不支持指针运算
	//println(b3)

}
