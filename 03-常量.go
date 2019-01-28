package main

import (
	"fmt"
	//"math"
)

// 常量必须是编译期可确定的数字、字符串、布尔值

const x, y int = 1, 2   // 多常量初始化
const s = "hello world" // 常量类型推断

// 常量组
const (
	a, b string = "hahahaha", "xixixiix"
	c, d int    = 3, 4
	e    bool   = true
)

// 枚举
// 关键字iota 定义常量组中从0开始按行计数的自增枚举值
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 类型声明
type (
	byte int8
	文本   string // 可以给类型别名
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
	// 类型转换 数字型可以相互转化 字符串 布尔 数字 不能相互转换

	var a float32 = 100.8
	fmt.Println(a)
	b := int32(a)
	fmt.Println(b)
	c := int8(a)
	fmt.Println(c)

}
