package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成四位随机数字
func CreatNum(randNum *int) {
	t := time.Now().UnixNano()
	rand.Seed(t)
	for {
		num := rand.Intn(10000)
		if num >= 1000 {
			*randNum = num
			break
		}
	}
}

// 将数字转换成长度为4的slice
func GetNum(s []int, randNum int) {

	s[0] = randNum / 1000 // 取千位数
	fmt.Println("千位数：", s[0])
	s[1] = randNum % 1000 / 100 // 取百位数
	fmt.Println("百位数：", s[1])
	s[2] = randNum % 100 / 10 // 取十位数
	fmt.Println("十位数：", s[2])
	s[3] = randNum % 10 // 取十位数
	fmt.Println("个位数：", s[3])
}

// 输入一个4位数并取出每一位
func OneGame(s []int) {
	var inum int
	for {
		fmt.Println("请输入一个4位数字：")
		fmt.Scan(&inum)
		if 999 < inum && inum < 10000 {
			break
		}
		fmt.Printf("输入的数字不符合要求，请重新输入 \n")
	}
	fmt.Printf("输入的数字是：%d", inum)
	k := make([]int, 4)
	GetNum(k, inum)
	fmt.Println("KeySlice", k)

	n := 0
	for i := 0; i < 4; i++ {
		if k[i] > s[i] {
			fmt.Printf("第%d位数大了一点\n", i+1)
		} else if k[i] < s[i] {
			fmt.Printf("第%d位数小了一点\n", i+1)
		} else {
			fmt.Printf("第%d位数猜对了\n", i+1)
			n++
		}
	}
	if n == 4 { // 全部才对
		fmt.Println("")

	}
}

func main() {
	var randNum int
	s := make([]int, 4)

	CreatNum(&randNum)
	fmt.Println("randNum:", randNum)

	GetNum(s, randNum)
	fmt.Println(s)
	OneGame(s)
}
