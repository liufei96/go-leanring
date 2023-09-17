package main

import "fmt"

func main() {

	printPyramid(8)

	printMulti(8)
}

// 打印金字塔的代码封装到函数中
func printPyramid(totalLevel int) {
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		// j 表示每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func printMulti(num int) {
	// 打印出九九乘法表
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Println()
	}
}
