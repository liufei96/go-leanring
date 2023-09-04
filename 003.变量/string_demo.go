package main

import "fmt"

func main() {
	stringDemo1()
	stringDemo2()
	stringDemo3()
	stringDemo4()
	stringDemo5()
}

func stringDemo1() {
	var address string = "浙江杭州 110 hello world!"
	fmt.Println(address)
}

// 字符串不可变
func stringDemo2() {
	var str = "hello"
	// str[0] = 'a' // 这里不能去修改str的内容
	fmt.Println(str)
}

func stringDemo3() {
	str2 := "abc\nabc"
	fmt.Println(str2)

	// 使用反引号 ``
	str3 := `
   func stringDemo1() {
		var address string = "浙江杭州 110 hello world!"
		fmt.Println(address)
	}
	`
	fmt.Println(str3)
}

// 字符串拼接
func stringDemo4() {
	var str = "hello" + "world"
	str += " haha!"
	fmt.Println(str)
}

// 当一行字符太长时，需要使用到多行字符串
func stringDemo5() {
	str4 := "hello" + "world" + "hello" + "world" + "hello" + "world" + "hello" + "world" +
		"hello" + "world"
	fmt.Println(str4)
}
