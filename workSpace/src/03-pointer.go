package main

import "fmt"

func testptr() *string {
	// 返回一个String类型的指针
	city := "北京"
	cityPtr := &city
	return cityPtr
}

func main() {

	res := testptr()
	fmt.Println("res is", *res)

	// C中指针访问结构体成员tree->next 在GO中只有tree.next

	// 01-直接申请指针
	name := "Lily"
	ptr := &name
	fmt.Println("name :", *ptr)
	fmt.Println("name ptr :", ptr)

	// 02-使用new关键字
	name2 := new(string)
	*name2 = "Bili"
	fmt.Println("name :", *name2)
	fmt.Println("name ptr :", name2)

	if res == nil {
		fmt.Println("res 为空")
	} else {
		fmt.Println("res 非空")
	}
}
