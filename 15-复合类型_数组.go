package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	//注意： 定义数组时 指定的数组元素个数必须是常量 不能是变量
	//n := 10
	const i = 10
	var num [i]int
	fmt.Println("num=", num)

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

	// 随机数
	// 设置种子，只需一次，如果种子参数一样，每次运行程序产生的随机数都一样(方法一)
	//rand.Seed(000)
	// （方法二）以当前系统时间戳作为种子参数
	t := time.Now().UnixNano()
	rand.Seed(t)
	for i := 0; i < 3; i++ {
		fmt.Println("rand1=", rand.Int())     // 随机很大的数 没限长度范围
		fmt.Println("rand2=", rand.Intn(100)) // 限制在100以内都数
	}

}
