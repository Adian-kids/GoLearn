package main

import (
	"fmt"
	. "importTest/sub"
	// SUB "importTest/sub"
	// "importTest/sub"
)

func main() {
	//res := SUB.Sub(20, 7)
	//res := sub.Sub(20, 7)

	res := Sub(20, 7)
	fmt.Println("sub.Sub(5,2) = ", res)
}
