package main

import (
	"fmt"
	util "go-leanring/006.func_package_error/utils"
	"go-leanring/006.func_package_error/递归"
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

	// 调用getSumAndSub
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("sum=%v, sub=%v", res1, res2)

	递归.Test1(4)
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

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
