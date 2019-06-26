package main

import "fmt"

//　接口 一个接口类型总是代表着某一种类型(既所有实现它的类型)
// 接口定义一个对象的行为，接口只指定了对象应该做什么，至于如何实现这个行为，则由对象本身去确定
// 所谓实现一个接口中的方法是指，具有与该方法相同的声明并且添加了实现部分（由花括号包裹的若干条语句）
type Animal interface {
	Grow()
	Move(new string) (old string)
}

type Cat struct {
	name   string
	age    uint8
	status string
}

type Human struct {
	name   string
	Gender string
	Age    uint8
}

// 根据公司员工的个人薪资，计算公司的总支出
type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay * p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// otalExpense 方法体现出了接口的妙用。该方法接收一个 SalaryCalculator 接口的切片（[]SalaryCalculator）作为参数。
func totalExpense(s [3]SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("总支出：%d \n", expense)
}

// 接口的内部表示
type Test interface {
	Teacher()
}

type MyFloat float64

func (m MyFloat) Teacher() {
	fmt.Println("m:", m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v \n", t, t)
}

// 空接口
func describe1(i interface{}) {
	fmt.Printf("Type = %T, value = %v \n", i, i)
}

// 类型断言
func assert(i interface{}) {
	s, ok := i.(int) // 从i中获取底层int值
	fmt.Printf("类型：%T, err:%v \n", s, ok)
}

// 类型选择
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("string 类型：%s \n", i.(string))
	case int:
		fmt.Printf("int 类型：%d \n", i.(int))
	default:
		fmt.Printf("unkown type \n")
	}
}

func main() {
	p := Human{"Robert", "Mala", 20}
	// 使用类型转换表达式把一个*Person类型转换成空接口类型的值：
	v := interface{}(&p)
	fmt.Println(v)
	myCat := Cat{"Little C", 2, "In the house"}

	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v \n", ok, animal)

	p1 := Permanent{1, 5000, 20}
	p2 := Permanent{2, 6000, 30}
	c1 := Contract{3, 3000}
	employee := [3]SalaryCalculator{p1, p2, c1}
	totalExpense(employee)

	//  定义Test接口类型的接口
	// MyFloat类型实现类该接口
	var t Test = MyFloat(89.7)
	//f := MyFloat(89.7)
	//t = f
	describe(t)
	t.Teacher()

	// 空接口
	// 没有包含方法的接口称为空接口，由于空接口没有方法，因此所有类型都能实现类空接口
	s := "Hello World" // string类型实现空接口
	describe1(s)
	i := 55 // int类型实现空接口
	describe1(i)
	strt := struct {
		name string
	}{"jack"}
	describe1(strt) // 结构体类型实现空接口

	// 类型断言
	var s1 interface{} = 56            // 类型为int不会报错
	var s2 interface{} = "Steven Paul" // 类型为string不会报错 用 v, ok := i.(T)不会报错
	assert(s1)
	assert(s2)

	// 类型选择
	findType("Naveen")
	findType(77)
	findType(99.44)

}
