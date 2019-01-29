package main

import "fmt"

func main() {
	//num := 2
	//switch num {  // switch 后面写的是变量本身 而不是条件 注意跟if有区别
	switch num := 20; num { // switch 后面支持一个初始化变量
	case 10:
		fmt.Println("按下的是1楼")
		//break   // 跳出switch语句， 默认不写 就包含含来break
		fallthrough // 强制执行case后面都流程都执行 不跳出switch语句
	case 20:
		fmt.Println("按下的是2楼")
	case 30:
		fmt.Println("按下的是3楼")
	default:
		fmt.Println("按下的是0楼") // 条件都不满足 默认值
	}

	score := 80
	switch { // switch 后面可以没有条件
	case score > 90: // case 后面可以放条件
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	}
}
