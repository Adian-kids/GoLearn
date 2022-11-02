package main

import "fmt"

func main() {
	// 01-定义
	name := "Adian\n"
	// 如果需要换行，则使用反引号``
	usage := `Hello
				I am a Usage
				Have fun`
	fmt.Println(name, usage)

	// 02-求长度以及访问
	length1 := len(usage)
	fmt.Println(length1)
	// 逐个输出
	for i := 0; i < len(name); i++ {
		fmt.Printf("i : %d , v ： %c\n", i, name[i])
	}
	// 更改内容
	name = "b"

	// 03-拼接
	i, j := "Hello", "World"
	println("i+j = ", i+j)

}
