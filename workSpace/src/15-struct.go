package main

import "fmt"

type People struct {
	name string
	age  int
}

func main() {
	bob := People{age: 15, name: "bob"}
	fmt.Println("his name :", bob.name, "\nhis age :", bob.age)
}
