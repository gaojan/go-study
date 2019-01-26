package main

import "fmt"

//var a = [...]int{2,7,6,9,0,5,4,3,1,8}

func main() {
	//var a [2]int  // 定义数据  int代表数组的长度
	//var b =  []int{1,2,3,4,9}  //不限定长度可以不指定，如果指定 则长度不能超过指定的长度，不满足指定长度自动补空值
	//var c = [9]string{"d", "c"}
	//var d = [20]int{19: 19}  //将下标为19的元素赋值为2
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//
	//i := [2]int{1,2}
	//f :=[2]int{1,5}
	//fmt.Println(i==f)
	//n := [...][3]int{
	//	{1,2,3},
	//	{4,5,6},
	//}
	//fmt.Println(n)
	// 数组的指针和指针数据
	x, y := 1, 2
	n, m, l := 3, 4, 5
	arr := [...]int{5: 2}
	// 数组的指针
	var pf *[6]int = &arr // pf的指针 指向arr的地址 &取地址符号

	// 指针数组
	pfArr := [...]*int{&x, &y} //[0xc00001a070 0xc00001a078] 保存了两个指针地址的数组
	abc := [...]*int{&n, &m, &l}
	fmt.Println(arr)
	fmt.Println(pf)
	fmt.Println(pfArr)
	fmt.Println(abc)

	// 冒泡排序
	//num := len(a)
	//for i := 0; i < num; i++{
	//	for j := i+1; j < num; j++{
	//		if a[i] < a[j]{
	//			tmp := a[i]  // 交换位置
	//			a[i] = a[j]
	//			a[j] = tmp
	//		}
	//	}
	//}
	//fmt.Println(a)

}
