package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 写文件
func WriterFile(path string) {
	// 新建文件 并打开文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 使用完毕需要关闭文件
	defer f.Close()
	str := "hello go hello python hello java----++++"
	f.WriteString(str)
}

// 读文件
func ReadFile(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭文件
	defer f.Close()
	data := make([]byte, 1024) // 2k大小
	// n代表从文件读取内容的长度
	i, err1 := f.Read(data)
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err:", err1)
		return
	}
	fmt.Printf("文件长度为：%d, 文件内容为：%v \n", i, string(data[:i]))
}

// 每次读取一行
func ReadBufioFile(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭文件
	defer f.Close()
	// 新建一个缓冲区，把内容放进缓冲区
	r := bufio.NewReader(f)
	for {
		// 遇到'\n'结束文件读取
		data1, err2 := r.ReadBytes('\n')
		if err2 != nil {
			// 文件已经读取到最后面了， 但是'\也被读取进去' 最后一行不被读取
			if err2 == io.EOF {
				break
			}
			fmt.Println("err2:", err2)
			return
		}
		fmt.Println(string(data1))
	}
}

func main() {
	// 关闭后无法使用打印输出
	//os.Stdout.Close()
	// 标准设备文件，默认已经打开，用户可以直接使用
	os.Stdout.WriteString("are you ok \n")
	fmt.Println("hello world")

	// 写入文件
	path := "./demo.txt"
	WriterFile(path)

	// 读取文件
	path1 := "./demo1.txt"
	ReadFile(path1)
	fmt.Println("--------")

	// 一行一行读取
	ReadBufioFile(path1)

}
