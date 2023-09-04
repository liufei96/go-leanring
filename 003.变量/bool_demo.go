package main

import (
	"fmt"
	"unsafe"
)

func main() {
	boolDemo1()
}

func boolDemo1() {
	var b = false
	fmt.Println("b=", b) // b= false
	// 注意事项
	// 1. bool 类型占用存储空间是 1字节
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b)) // b 的占用空间 = 1
}
