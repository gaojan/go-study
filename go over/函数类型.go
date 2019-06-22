package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

// go语言中函数是一等类型
// 使用type定义函数

// 小例子：
// 员工ID生成器 函数类型
type EmployeeIdGenerator func(company string, department string, sn uint32) string

// 默认公司名称
var company = "xrt"

// 默认序列号
var sn uint32

// 生成员工id
func generateId(generator EmployeeIdGenerator, department string) (string, bool) {
	// 如果生成器不可用直接返回
	if generator == nil {
		return "", false
	}
	// 使用代码包 sync/atomic中提供的原子操作函数可以保证并发安全
	newSn := atomic.AddUint32(&sn, 1)
	fmt.Println("newSn:", newSn)
	return generator(company, department, newSn), true
}

// 字符串类型和数值类型不可直接拼接，需要转换
func appendSn(firstPart string, sn uint32) string {
	return firstPart + strconv.FormatUint(uint64(sn), 10)
}

func main() {
	var generator EmployeeIdGenerator
	//sn1 = appendSn("-", 2)

	generator = func(company string, department string, sn uint32) string {
		str := company + "-" + department + appendSn("-", sn)
		return str
	}

	fmt.Println(generateId(generator, "jsb"))
}
