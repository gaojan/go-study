package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Contains 判断是否包含某子字符串，包含返回true 不包含返回false
	fmt.Println(strings.Contains("helloworld", "hello"))

	// Join 组合拼接字符串 第一个参数是切片或数组 第二个参数是分割符
	s := []string{"abc", "def", "efg"}
	fmt.Println(strings.Join(s, ","))

	// Index 查找子字符串索引位置  存在返回索引位置，不存在返回-1
	fmt.Println(strings.Index("helloworld", "world"))

	// Repeat 重复字符串 n表示重复次数，小于0表示全部替换
	fmt.Println(strings.Repeat("hellogo", 3)) //hellogohellogohellogo

	// Replace 替换前n个字符串 -1 代表所有字符串  n代表前几个
	fmt.Println(strings.Replace("hello", "l", "o", -1))

	// Split 分割字符串，返回slice，以指定的分隔符分割
	fmt.Println(strings.Split("hello,world", ","))

	// Trim  去掉两头空格
	fmt.Println(strings.Trim("  are you ok  ", " "))

	// Fields 去掉空格，把元素放入切片中
	s1 := strings.Fields("   are you ok  ")
	fmt.Println(s1)

	//Append slice转换字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	// 第一个为slice 第二个数为要追加都的数，第三个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	fmt.Printf("%v\n", slice)
	slice = strconv.AppendQuote(slice, "abcdef")
	fmt.Printf("%s \n", string(slice))

	// bool类型转换字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println(str)

	// float类型转字符串
	// 'f' 指打印格式，以小数方式，prec -1指小数点位数（紧缩模式），bitsize 64 以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println(str)

	// Itoa 整型转换字符串
	var n int
	n = 9999
	str = strconv.Itoa(n)
	fmt.Printf("%s \n", str)

	// 字符串转bool
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag=", flag)
	} else {
		fmt.Println("err:", err)
	}

	// Atoi 字符串转整型
	a, _ := strconv.Atoi("8888")
	fmt.Printf("%d \n", a)

}
