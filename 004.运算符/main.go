package main

import "fmt"

func main() {
	demo1()
}

func demo1() {
	fmt.Println(10 / 4) // 2

	var n1 float32 = 10 / 4
	fmt.Println(n1) // 2

	// 如果我们西王保留小水部分，则需要浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2) // 2.5
}
