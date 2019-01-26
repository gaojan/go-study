package main

import "fmt"

func main() {
	//for a := 1; a<10; a++{  //循环语句可以跟初始化、条件语句写在一起 ， 不需要写break 会自动终止 如果需要继续下一个case则需要写fallthrough语句
	//	fmt.Println(a)
	//}
	//switch 选择语句
	//b := 6  // 如果这里不写变量 则可以在case中写表达式
	//switch {
	//case b >= 1:
	//	fmt.Println("b=1")
	//	fallthrough
	//case b >= 2:
	//	fmt.Println("b=2")
	//default :
	//	fmt.Println("None")
	//}
	// 跳转语句 break、continue、goto  三个语句都可以配合标签 goto是调整执行位置
	//标签名区分大小写，若不使用标签会造成编译错误
	//AA:
	//	for {
	//		for a := 1; a < 10; a++{
	//			if a > 5 {
	//				//break AA   // 满足条件到这里终止
	//				//goto AA  // goto 调整执行位置 又调整到AA，会一直循环下去
	//				continue // 可以不使用标签
	//			}
	//			fmt.Println(a)
	//		}
	//	}
	//AA:  //标签放这里 会避免死循环

BB:
	for a := 1; a < 10; a++ {
		for {
			fmt.Println(a)
			//continue BB
			goto BB // 会一直执行初始化值
		}
		fmt.Println("OK")
	}
}
