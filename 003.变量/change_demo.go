package main

import "fmt"

func main() {
	changeDemo1()
	changeDemo2()
	changeDemo3()
}

// 类型转换
func changeDemo1() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v, n1=%v, n2=%v, n3=%v", i, n1, n2, n3)
}

func changeDemo2() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v, n1=%v, n2=%v, n3=%v", i, n1, n2, n3)
	// 被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化！
	fmt.Printf("i type is %T\n", i) // i type is int32
}

// 在转换中，比如将 int64 转成 int8 【-128~127】,编译时不会报错
// 只是转换的结果是按溢出处理，和我们希望的结果不一样
func changeDemo3() {
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2) // num2= 63
}
