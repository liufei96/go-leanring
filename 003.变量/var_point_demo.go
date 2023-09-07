package main

import "fmt"

func main() {
	varPointDemo1()
	varPointDemo2()
}

func varPointDemo1() {
	// 基本数据类型在内存的布局
	var i int = 10
	// i 的地址是什么，&i
	fmt.Println("i的地址=", &i)

	// 下面的 var ptr *int = &i
	// 1. ptr 是一个指针变量
	// 2. ptr 的类型 *int
	// 3. ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v", &ptr)
	fmt.Printf("ptr 指向的值=%v", *ptr)
}

func varPointDemo2() {
	var num int = 10
	fmt.Printf("num address=%v\n", num)

	var ptr *int
	ptr = &num
	*ptr = 11 // 这里修改时，会到num的值变化
	fmt.Println("num =", num)
}
