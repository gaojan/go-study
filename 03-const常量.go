package main

import (
	"fmt"
	//"math"
)

// 常量必须是编译期可确定的数字、字符串、布尔值 程序运行期间不可改变的值

const x, y int = 1, 2   // 多常量初始化
const s = "hello world" // 常量类型推断

// 常量组
const (
	a, b string = "hahahaha", "xixixiix"
	c, d int    = 3, 4
	e    bool   = true
)

// 枚举
// 关键字iota 定义常量组中从0开始按行计数的自增枚举值 每个一行 自动累加1
// 遇到const 重置为0
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	aa = iota
	bb
	cc
)

func main() {
	//println(x, y)
	//println(s)
	//println(a, b)
	//println(c, d)
	//println(e)
	//fmt.Println(math.MaxInt32)
	//var  b 文本
	//var b 文本 = "中文类型名"
	//fmt.Println(b)

	fmt.Println(aa, bb, cc)

}
