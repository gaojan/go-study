// 每个go文件开头都拥有一个package声明，表示源码文件所属代码包
// 必须要有main的package包，且必须要有main()函数
// 一个go文件中 必须要有main 函数，go文件的入口，否则不能执行
package main

import (
	//o "fmt"  // 包的别名
	//_"fmt"   // _ 表示使用空白标示符 可以不用这个包 但程序不报错
	"fmt"
)

// 类型声明
type (
	byte int8
	文本   string // 可以给类型别名
)

func main() {
	//o.Print("hello world")
	//o.Printf("""")  // 打印格式化 使用
	//o.Fprintln("")  // 打印换行使用  ln代表换行
	// 类型转换 数字型可以相互转化 字符串 布尔 数字 不能相互转换
	// 字符转整型
	var aaa byte
	aaa = 'a'
	var tt int
	tt = int(aaa) // 把aaa的值取出来后 转换成int再给tt赋值
	fmt.Println(tt)

	// 变量的输入 类似input
	var h int
	//fmt.Scanf("%d:",&h)  // 加上&7
	fmt.Scan(&h) // 更简洁的写法

	var a float32 = 100.8
	fmt.Println(a)
	b := int32(a)
	fmt.Println(b)
	c := int8(a)
	fmt.Println(c)

	var ch byte
	var str string
	ch = 'a'
	str = "helloworld"
	fmt.Println(ch)                                      // 97  单引号 字符往往都是一个字符，转义字符除外'\n'
	fmt.Println(str)                                     // 双引号  字符串都是隐藏了一个结束符 '\0'
	fmt.Printf("str[0]:%c, str[1]:%c\n", str[0], str[1]) //类似于字符串的切片

	//复数类型
	var t complex64
	t = 2.1 + 3.14i
	fmt.Println(t)
	// 可以通过内建函数取实部和虚部
	fmt.Println("real:", real(t), "imag", imag(t))

	// 格式化输出 常用的
	// %c 字符型
	// %d 十进制数值
	// %b 二进制整数值
	// %t 以true或false输出的布尔值
	// %T 值的类型
	// %v 内置或自定义类型的值  万能格式 自动配置值
	// %+v 显示更详细
}
