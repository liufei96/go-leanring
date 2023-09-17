package main

import (
	"fmt"
)

func main() {
	test(4)
	test2(4)
	res := fbn(10)
	fmt.Println("res=", res)

	res1 := f(2)
	fmt.Println("res1=", res1)

	res2 := peach(1)
	fmt.Println("第1天桃子的数量=", res2)

	// n的值不会改变
	n := 10
	test3(n)
	fmt.Println("main() n=", n)

	// 可以传入变量地址
	num := 20
	test4(&num)
	fmt.Println("main() num=", num)

	a := getSum
	fmt.Printf("a的类型%T,getSum类型时%T\n", a, getSum)
	result := a(10, 20)
	fmt.Println("result=", result)
}

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fbn(n-1) + fbn(n-2)
}

func f(n int) int {
	if n == 1 {
		return 3
	}
	return 2*f(n-1) + 1
}

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("你输入的天数不对!!")
		return 0
	}
	if n == 10 {
		return 1
	}
	return (peach(n+1) + 1) * 2
}

func test3(n int) {
	n = n + 10
	fmt.Println("test3 n=", n)
}

func test4(n *int) {
	*n = *n + 20
	fmt.Println("test4 n=", *n) // 如果直接打印n，打印的是变量地址
}

// 在Go中，函数也是一种数据类型
// 可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}
