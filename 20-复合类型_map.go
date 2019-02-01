package main

import (
	"fmt"
	"sort"
)

// 类似于哈希表或者python中的字典 key-value
// map做函数参数，实际上所引用传递
func test(m3 map[int]string) {
	delete(m3, 4)
	fmt.Println("test: m3=", m3) // 里面修改删除参数外面也会影响
}

func main() {
	// map创建
	//var m map[int]string  // 定义一个m变量 m为字典
	var m1 = make(map[int]string) //跟上面效果一样
	m1[1] = "ok"
	a := m1[1]
	fmt.Println("m1:", m1)
	fmt.Println("a:", a)

	// 多级map嵌套  需要没级的map初始化
	var m2 map[string]map[string]string
	m2 = make(map[string]map[string]string) // 一级map初始化
	m2["1"] = make(map[string]string)       // 二级map初始化
	m2["haha"] = map[string]string{"123": "456"}
	m2["1"]["5"] = "hello"
	fmt.Println("m2:", m2)
	// map遍历
	for k, v := range m2 {
		fmt.Println("m2,k-v:", k, v)
	}
	// 如何判断一个key值是否存在
	// 第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在ok为true
	value, ok := m1[1]
	if ok == true {
		fmt.Println("m1[1]=", value)
	} else {
		fmt.Println("key不存在")
	}

	// 将map中的k，v进行交换，并排序
	p := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	var p1 = map[string]int{}
	var arr = make([]int, len(p))
	i := 0
	for k, v := range p {
		arr[i] = k
		p1[v] = k
		i++
	}
	sort.Ints(arr) // 只能对slice排序 字典本身无序
	fmt.Printf("arr--%v \n", arr)
	fmt.Printf("p--%v \n", p)
	fmt.Printf("p1--%v \n", p1)

	// map删除
	m := map[int]string{1: "mike", 2: "yoyo", 3: "gogo"}
	fmt.Println("m=", m)
	delete(m, 1) // delete(mapName, Key) 删除key为1的内容
	fmt.Println("m=", m)

	// map做函数引用
	m3 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	test(m3)
	fmt.Println("main: m3=", m3)

}
