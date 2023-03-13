package main

import "fmt"

func main() {
	// goto LABEL_1
	// break LABEL_2
	// continue LABEL_3
Label1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break Label1
			}
			fmt.Println("i = ", i, "j = ", j)
		}
	}
}
