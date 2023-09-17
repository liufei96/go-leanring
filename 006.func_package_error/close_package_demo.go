package main

import (
	"fmt"
	"strings"
)

func main() {
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16

	// 测试makeSuffix 的使用
	// 返回一个闭包

	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))
	fmt.Println("文件名处理后=", f2("bird.jpg"))
}

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(i int) int {
		n = n + i
		str += string(36) // 36 => $
		fmt.Println("str=", str)
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(file string) string {
		if !strings.HasSuffix(file, ".jpg") {
			file = file + ".jpg"
		}
		return file
	}
}
