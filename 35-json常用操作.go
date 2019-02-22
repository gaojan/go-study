package main

import (
	"encoding/json"
	"fmt"
)

/*

{
"company": "itcast",
"subjects":[
	"Go",
	"c++"
],
"isok": true,
"price": 666.666
}

*/

// 方法一：通过结构体生成json
//type IT struct {
//	//json成员变量名首字母必须大写
//	Company string
//	Subjects []string
//	Isok bool
//	Price float64
//}

type IT struct {
	Company  string   `json:"company"`  // 二次编码 可以全部转为小写
	Subjects []string `json:"subjects"` // `json:"-"` 此字段不会输出到屏幕
	Isok     bool     `json:"isok"`     // `json:"string"` bool或int、float 转换成字符串
	Price    float64  `json:"price"`
}

func main() {
	fmt.Println("-----方法一：通过结构体生成json-----")
	// 方法一：通过结构体生成json
	i := IT{"itcast", []string{"GO", "C++"}, true, 666.666}
	// 编码， 根据内容生成json文本
	//s,_ := json.Marshal(i)
	//fmt.Printf("%s \n", string(s))
	//s, err := json.Marshal(i)  // 非格式编码
	s, err := json.MarshalIndent(i, "", "") // 格式化编码
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%s \n", string(s))
	}
	fmt.Println("-----方法二： 通过map生成结构体-----")

	// 方法二： 通过map生成结构体
	m := make(map[string]interface{})
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++"}
	m["isok"] = true
	m["price"] = 666.666

	j, err := json.MarshalIndent(m, "", "")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%s \n", j)
	}
	fmt.Println("--------反向解析json到结构体----------")

	// 反向解析json到结构体
	jsonBuf := `{
	"company": "itcast",
	"isok": true,
	"price": 666.666,
	"subjects": [
		"Go",
		"C++"
	]
}`
	var tmp IT
	err = json.Unmarshal([]byte(jsonBuf), &tmp) // 第二个参数要传地址变量
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v \n", tmp) // % +v 格式化输出更详细的内容值

	// json只解析其中某一个字段
	type IT2 struct {
		Subjects []string `json:"subjects"`
	}
	var t IT2
	err = json.Unmarshal([]byte(jsonBuf), &t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v \n", t)
	fmt.Println("-------反向解析json到map------")

	// 反向解析json到map
	m1 := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBuf), &m1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v \n", m1)
	fmt.Println("-----json只解析其中某一个字段-----")

	// 此方法直接转换会报错
	//var st []string
	//st = m["subjects"]
	//err = json.Unmarshal([]byte(jsonBuf), &st)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%+v", st)

	// 类型断言
	var st []string
	var n float64
	var s1 string
	for k, v := range m {
		//fmt.Printf("%v=======%v \n", k, v)
		switch data := v.(type) {
		case []string:
			st = data
			fmt.Printf("%v-----%v \n", k, st)
		case float64:
			n = data
			fmt.Printf("%v-----%v \n", k, n)
		case string:
			s1 = data
			fmt.Printf("%v-----%v \n", k, s1)
		}
	}
}
