package main

import "fmt"

func main() {
	city_all := [5]string{"北京", "上海", "广州", "天津", "唐山"}
	fmt.Println("city_all 初始定义: ", city_all)

	// 01-切片操作,取下标为0，1，2的加入新的数组city_1
	city_part_1 := city_all[0:3]
	fmt.Println("city_part_1 初始定义: ", city_part_1)

	// 02-修改切片中的值会改变原数组的值
	city_part_1[1] = "周杰伦"
	fmt.Println("city_all 修改后: ", city_all)
	fmt.Println("city_part_1 修改后 : ", city_part_1)

	// 03-创建指定容量cap的切片
	test_cap := make([]string, 10, 20)
	println("test_cap len :", len(test_cap), "cap :", cap(test_cap))

	// 04-切片独立于原数组,如果全部copy，则需要使用[:]将数组转为切片
	city_copy := make([]string, 3, 3)
	copy(city_copy, city_all[0:3])
	fmt.Println("city_copy :", city_copy)

}
