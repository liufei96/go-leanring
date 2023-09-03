package main

import (
	"fmt"
	"unsafe"
)

func main() {
	intDemo1()
	intDemo2()
	intDemo3()
	fmt.Println("")
	intDemo4()
}

func intDemo1() {
	var i int = 10
	fmt.Println("i=", i)

	var j int8 = 127
	fmt.Println("j=", j)

	var k uint8 = 255
	fmt.Println("k=", k)

	var b uint = 1
	fmt.Println("b=", b)

	var c byte = 255
	fmt.Println("c=", c)
}

func intDemo2() {
	// 整型的使用细节
	var n1 = 10 // ? n1 是什么类型
	// 这里我们给介绍一下如何查看某个变量的数据类型
	fmt.Printf("n1 的类型 %T \n", n1) // n1 的类型 int
}

// 如何在程序查看某个变量的字节大小和数据类型 （使用较多）
func intDemo3() {
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T  n2占用的字节数是 %d ", n2, unsafe.Sizeof(n2)) // n2 的类型 int64  n2占用的字节数是 8
}

// Golang 程序中整型变量在使用时，遵守保小不保大的原则，即：在保证程序正确运行下，尽量使用占用空间小的数据类型。【如：年龄】
func intDemo4() {
	var age byte = 90
	fmt.Printf("age 使用 %T 类型", age)
}
