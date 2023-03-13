package main

import "fmt"

// 定义常量,模拟一周的枚举
const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	SateDay
	Sunday
	M, N = iota, iota
)

func main() {
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(SateDay)
	fmt.Println(Sunday)

}
