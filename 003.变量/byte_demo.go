package main

import "fmt"

func main() {
	byteDemo1()
	byteDemo2()
}

func byteDemo1() {
	var c1 byte = 'a'
	var c2 byte = '0' // 字符的0

	// 当我们直接输出byte的值，就是输出了对哦应字符的码值
	fmt.Println("c1=", c1) // 97
	fmt.Println("c2=", c2) // 48
	// 如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c  c2=%c", c1, c2) // c1=a  c2=0

	// var c3 byte = '北'  // overflow 溢出
	var c3 int = '北'
	fmt.Printf("c3=%c  c3对应码值=%d", c3, c3) // c3=北  c3对应码值=21271
}

func byteDemo2() {
	var c4 int = 22269
	fmt.Printf("c4=%c \n", c4) // c4=国

	var n1 = 10 + 'a'
	fmt.Println("n1=", n1) // n1= 107
}
