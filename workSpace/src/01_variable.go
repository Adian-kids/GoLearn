package main

import "fmt"

func main() {
	// 定义变量 var
	// 定义常量 const

	// 01-先定义变量再赋值
	// var [varname] [datatype]
	var age int
	var name string
	age = 3
	name = "Adian"
	// printf格式化输出或者println输出整行
	fmt.Println("name is", name, "and age is", age)
	fmt.Printf("name is %s and age is %d\n", name, age)

	// 02-定义时直接赋值，类型可以省略
	var gender = "man"
	fmt.Println("gender is", gender)

	// 03-定义直接赋值，自动识别类型
	height := 172.3
	fmt.Println("height is", height)

	// 04-平行赋值
	i, j := 10, 20
	fmt.Println("i=", i, "j=", j)

}
