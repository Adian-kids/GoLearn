package main

import (
	"fmt"
	"os"
)

// 从命令行输入参数，在switch中进行判断
// C : argc **argv
// Go: os.Args,收到一个字符串切片
func main() {
	cmds := os.Args
	//for key, cmd := range cmds {
	//	fmt.Println("key :", key, " cmd :", cmd)
	//}
	switch cmds[1] {
	case "help":
		fmt.Println("help info")
	case "test":
		fmt.Println("test info")
	default:
		fmt.Println("default info")
	}

}
