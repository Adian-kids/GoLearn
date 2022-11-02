package main

import "fmt"

func main() {
	// 01-定义长度为10的数组
	nums1 := [10]int{1, 2, 3, 4, 5}
	var nums2 = [10]int{6, 7, 8, 9, 0}
	var nums3 [10]int = [10]int{1, 2, 3, 4}
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(nums3)

	// 02-访问和遍历
	// 逐个输出
	for i := 0; i < len(nums3); i++ {
		fmt.Printf("i : %d , v ： %d\n", i, nums3[i])
	}
	// 可以和Python一样for range
	// key是数组下标，value是值
	for key, value := range nums3 {
		fmt.Println("key : ", key, " value : ", value)
	}
	// 如果想忽略key或者value,可以使用_
	for _, value := range nums3 {
		fmt.Println("key : None", " value : ", value)
	}
	// 如果两个都忽略，则需不需要:
	for _, _ = range nums3 {
		fmt.Println("key : None", " value : ")
	}

	// 修改对应的内容
	nums1[1] = 999
	fmt.Println(nums1)
}
