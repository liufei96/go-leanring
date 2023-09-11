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
