package main

import (
	"fmt"
	"go-leanring/006.func_package_error/utils"
)

var age = init_test()

func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...")
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}

func init_test() int {
	fmt.Println("init_test()") // i
	return 90
}
