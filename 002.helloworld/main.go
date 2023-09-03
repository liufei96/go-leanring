package main

import "fmt"

func main() {
	test1()
	test2()
}

// 1. 转义字符的使用
func test1() {
	// 加个tab
	fmt.Println("tom\tjack")
	// 换行
	fmt.Println("hello\nworld")
	// \\
	fmt.Println("F:\\git_project\\go-leanring")
	// \"
	fmt.Println("\"你好\"")
	// 回车 \r，从当前行的最前面开始输出，覆盖掉以前的内容

	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
}

// 2. 注释的内容
func test2() {
	//fmt.Println("单行注释")
	/*	fmt.Println("多行注释")
		fmt.Println("多行注释")*/
}
