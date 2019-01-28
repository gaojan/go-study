// 每个go文件开头都拥有一个package声明，表示源码文件所属代码包
// 必须要有main的package包，且必须要有main()函数
package main

import (
	o "fmt"
)

func main() {
	o.Print("hello world")
}
