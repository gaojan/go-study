package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

// 将切片中的所有整数乘以5，并返回接口
func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	s1 := student{
		"Naveen",
		"Ramanathan",
		"A",
		"India",
	}
	s2 := student{
		"Samuel",
		"Johnson",
		"B",
		"USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
	// 假设我们想要查找所有来自印度的学生，通过修改传递给filter的函数参数
	c := filter(s, func(s student) bool {
		if s.country == "India" {
			return true
		}
		return false
	})
	fmt.Println(c)
	// 上述用途 有点像python中的装饰器

	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)

}
