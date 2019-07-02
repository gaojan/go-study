package main

import "fmt"

// 使用接口实现多态
// 一个类型如果定义了接口的所有方法，那它就隐式地实现了该接口。
// 通过一个程序理解go的多态，计算一个组织机构的净收益，组织所获得的收入来源于两个项目：fixed billing 和 time and material。
// 该组织的净收益等于这两个项目的收入总和。

// 定义一个收入接口
type Income interface {
	calculate() int
	source() string
}

// 项目FixedBillin
type FixedBilling struct {
	projectName  string
	biddedAmount int
}

// TimeAndMaterial 用于表示项目Time and Material
type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

// 给结构体定义方法
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

// 此函数用来计算并打印总收入
func calculateNetIncome(ic []Income) {
	netincome := 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d \n", netincome)
}

func main() {
	/*
		创建了三个项目，有两个是 FixedBilling 类型，一个是 TimeAndMaterial 类型。接着我们创建了一个 Income 类型的切片，
		存放了这三个项目。由于这三个项目都实现了 Interface 接口，因此可以把这三个项目放入 Income 切片。最后我们将该切片作为参数，
		调用了 calculateNetIncome 函数，显示了项目不同的收益和收入来源。
	*/
	project1 := FixedBilling{"project 1", 5000}
	project2 := FixedBilling{"project 2", 10000}
	project3 := TimeAndMaterial{"project 3", 160, 25}
	bannerAd := Advertisement{"Banner Ad", 2, 500}
	popupAd := Advertisement{"Popup Ad", 5, 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)

	/*
		尽管我们新增了收益流，但却完全没有修改 calculateNetIncome 函数。这就是多态带来的好处。
	*/

}
