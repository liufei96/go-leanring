package main

import (
	"fmt"
	"os"
)

func main() {
	res := deferSum(10, 20)
	fmt.Println("res=", res)
}

func deferSum(n1 int, n2 int) int {
	// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)

	// 增加一句话
	n1++
	n2++

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func testFile(filename string) {
	// 关闭文件资源
	file, _ := os.Open(filename)
	defer file.Close()
	// 其它代码
}
