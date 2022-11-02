package main

import "fmt"

func main() {
	// 01-字典的定义
	// 学号==>姓名
	var idNames map[int]string // 此时map不可以直接使用
	// 给map分配空间
	idNames = make(map[int]string, 10) // 也可以不指定长度
	// 赋值
	idNames[0] = "BigB"
	idNames[1] = "BigC"
	// 直接输出
	fmt.Println(idNames)
	// 遍历输出
	for key, value := range idNames {
		println("key :", key, "value :", value)
	}
	// 定义时直接分配空间
	idNames1 := make(map[int]string, 10)
	println(idNames1)

	// 02-key的存在性判断
	// 再map中所有的key都是合法的，如果并不存在该key值，则返回nil,所以不能通过value来判断key的存在性
	// 因为可能value本身就是0,nil,false这些值
	value, ok := idNames[3]
	if ok {
		fmt.Println("存在value ：", value)
	} else {
		fmt.Println("不存在value")
	}

	// 03-通过key删除map中的元素
	delete(idNames, 0)

}
