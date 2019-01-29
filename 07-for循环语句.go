package main

import "fmt"

// 跳转语句 break、continue、goto  三个语句都可以配合标签 goto是调整执行位置

func main() {
	//for a := 1; a<10; a++{  //初始化条件；判断条件；条件变化， 不需要写break 会自动终止 如果需要继续下一个case则需要写fallthrough语句
	//	fmt.Println(a)
	//}

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

	//BB:
	//	for a := 1; a < 10; a++ {
	//		for {
	//			fmt.Println(a)
	//			//continue BB
	//			goto BB // 会一直执行初始化值  可以用在任何地方，但是不能夸函数使用 可以跳转到标签位置
	//		}
	//		fmt.Println("OK")
	//	}

	// range 迭代器 默认返回两个元素  一个是元素位置 一个是元素本身
	str := "hello world"
	for i, data := range str {
		fmt.Println(i, "---", string(data))
	}
}
