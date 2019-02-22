package main

import "fmt"
import "errors"

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = a / b
	}
	return
}

func main() {
	// 实现error接口方法一
	err1 := fmt.Errorf("%s", "This is error")
	fmt.Println("err1:", err1)
	// 实现error接口方法二
	err2 := errors.New("This is error2")
	fmt.Println("err2:", err2)

	result, err := MyDiv(1, 2)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("result=", result)
	}
}
