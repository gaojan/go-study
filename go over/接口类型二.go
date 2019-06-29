package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

type Address struct {
	state   string
	country string
}

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

// 接口的嵌套
type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeave  int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d \n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeave - e.leavesTaken
}

// 使用值接收者实现
func (p Person) Describe() {
	fmt.Printf("%s is %d years old \n", p.name, p.age)
}

// 使用指针接收者实现
func (a *Address) Describe() {
	fmt.Printf("state: %s, country: %s \n", a.state, a.country)
}

// 接口的零值
type Des interface {
	Desc()
}

func main() {
	var d1 Describer
	p1 := Person{"Sam", 30}
	d1 = p1
	d1.Describe() // 值接收

	p2 := Person{"James", 20}
	d1 = &p2
	d1.Describe() // 指针接收

	var d2 Describer
	a := Address{"ShenZhen", "China"}
	//d2 = a // 这是不合法的 报错
	d2 = &a
	d2.Describe()

	/*
		上面说明 1、使用值接受者声明的方法，既可以用值来调用，也能使用指针调用，
		不管是一个值还是一个可以解引用的指针，调用这样的方法都是合法的。
		2、使用指针接受者声明的方法，用一个指针或一个可取得地址的值来调用都是合法的
	*/

	// 实现多个接口
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeave:  30,
		leavesTaken: 5,
	}
	var s SalaryCalculator = e
	s.DisplaySalary()

	var l LeaveCalculator = e
	l.CalculateLeavesLeft()
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())

	// 接口的嵌套
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\n Leaves left =", empOp.CalculateLeavesLeft())

	// 接口的零值
	var des Des
	if des == nil {
		fmt.Printf("des is nil and has type %T value %v \n", des, des)
	}
	//des.Desc()
	//对于值为nil的接口， 其底层值和具体类型都为nil

}
