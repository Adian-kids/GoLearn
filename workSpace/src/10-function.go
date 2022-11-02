package main

func test(a int, b int) (sum int, str string) {
	// 直接把返回变量的变量名指定
	sum = a + b
	str = "hello"
	return
}

func test1(a int, b int) (int, string) {
	// 多返回值
	return a + b, "hello"
}

func test2(a int, b int) int {
	// 单返回值
	return a
}
func test3(a, b int) int {
	// 省略变量类型
	return a
}

func main() {
	sum, str := test(1, 3)
	println("sum :", sum, "str :", str)
}
