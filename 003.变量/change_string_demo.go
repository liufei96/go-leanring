package main

import (
	"fmt"
	"strconv"
)

func main() {
	changeStringDemo1()
	changeStringDemo2()
	changeStringDemo3()
	changeStringDemo4()
}

func changeStringDemo1() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string // 空的str

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q\n", str, str) // str type string str = "99"

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %q\n", str, str) // str type string str = "23.456000"

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q\n", str, str) // str type string str = "true"

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str = %q\n", str, str) // str type string str = "h"
}

func changeStringDemo2() {
	var num5 int64 = 99
	// Itoa是FormatInt(i, 10) 的简写。
	str := strconv.Itoa(int(num5))
	fmt.Printf("str type %T str = %q\n", str, str) // str type string str = "99"
}

func changeStringDemo3() {
	var str string = "true"
	var b bool
	// 1. strconv.ParseBool(str) 会返回两个值 （value bool, err error）
	// 2. 因为我只想获取到value bool，不想获取err 所以我使用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "1234567"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	// 第二个参数是bitSize 表示f的来源类型(32: float32、64: float64)
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}

func changeStringDemo4() {
	var str4 string = "hello"
	var n3 int64 = 11
	// 第二个参数是base，表示进制的意思，2~32
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T n3=%v\n", n3, n3) // n3 type int64 n3=0
}
