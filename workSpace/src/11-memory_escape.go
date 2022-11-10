package main

import "fmt"

func testPtr() *string {
	// 返回一个String类型的指针
	city := "北京"
	cityPtr := &city
	return cityPtr
}

func main() {
	p1 := testPtr()
	fmt.Println("p1 :", *p1)
}
