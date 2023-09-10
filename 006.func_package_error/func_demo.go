package main

import (
	"fmt"
	util "go-leanring/006.func_package_error/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '-'
	fmt.Println("res=", util.Cal(n1, n2, operator))

	n3 := 10
	test(n3)

	fmt.Println("main() n1=", n1) // ? 输出结果=?

	sum := getSum(10, 20)
	fmt.Println("main sum = ", sum)
}

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1 = ", n1) // 输出结果
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum() sum = ", sum)
	return sum
}
