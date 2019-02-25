package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcFileName, dstFileName string) {
	// 只读方式打开源文件
	rf, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("打开源文件错误:", err1)
		return
	}
	// 新建目的文件
	wf, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("新建目标文件错误:", err2)
		return
	}
	// 关闭文件
	defer rf.Close()
	defer wf.Close()

	// 循环读取文件，读一点写一点
	buf := make([]byte, 1024*2) // 生成2k大小内存
	for {
		i, err := rf.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				break
			}
			fmt.Println("读取文件错误：", err)
		}
		wf.Write(buf[:i])
	}

}

func main() {
	list := os.Args
	srcFileName := list[1]
	dstFileName := list[2]

	// 验证源文件和目标文件参数
	if len(list) != 3 {
		fmt.Println("usage: srcFileName dstFile")
		return
	}

	// 源文件和目标文件不能一样
	if srcFileName == dstFileName {
		fmt.Println("目标文件名不能和源文件名相同")
		return
	}
	copyFile(srcFileName, dstFileName)

}
