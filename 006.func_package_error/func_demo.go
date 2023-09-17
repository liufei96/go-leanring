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

	// 调用getSumAndSub
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("sum=%v, sub=%v", res1, res2)

	fmt.Println("")

	demo()

	res3 := myFunc2(getSum, 500, 600)

	fmt.Println("res3=", res3)

	res4 := sum1(1, 2, 3, 4, 5)
	//res4 := sum1()
	fmt.Println("res4=", res4)

	res5 := sum2(1)
	fmt.Println("res5=", res5)
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

func demo() {
	type myInt int

	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1) // 这里依然需要显示转换，go认为myInt和int是两个类型
	fmt.Println("num1=", num1, "num2=", num2)
}

// 这时 myFun 就是 func(int, int) int 类型
type MyFuncType func(int, int) int

// 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFunc2(funcVar MyFuncType, num1 int, num2 int) int {
	return funcVar(num1, num2)
}

// 支持0到多个参数
func sum1(args ...int) int {
	res := 0
	for i := range args {
		res += args[i]
	}
	return res
}

// 支持1到多个参数
func sum2(n1 int, args ...int) int {
	res := n1
	for i := range args {
		res += args[i]
	}
	return res
}

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
