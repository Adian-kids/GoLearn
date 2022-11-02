package main

import "fmt"

func main() {
	// 01-不定长数组(切片slice)的定义
	// 定长数组
	names1 := [10]string{"北京", "上海", "广州", "深圳"}
	// 切片(不定长数组)
	names2 := []string{"北京", "上海", "广州"}
	fmt.Println(names1, names2)

	// 02-追加数据
	names2 = append(names2, "天津")
	fmt.Println(names2)

	// 02-切片不仅有len()长度，还有容量cap()的概念
	fmt.Println("names2 append之前长度len : ", len(names2), "容量cap : ", cap(names2))
	names2 = append(names2, "鹤岗")
	fmt.Println("names2 append之后长度len : ", len(names2), "容量cap : ", cap(names2))

	// 03-容量申请测试
	testcap := []int{}
	for i := 0; i < 20; i++ {
		fmt.Println("testcap append之后长度len : ", len(testcap), "容量cap : ", cap(testcap))
		testcap = append(testcap, i)
	}
}
