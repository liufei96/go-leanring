package main

import "fmt"

var (
	Func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	anonDemo1()
	anonDemo2()

	res4 := Func1(4, 9)
	fmt.Println("res4=", res4)
}

func anonDemo1() {
	// 在定义匿名函数时就直接调用，这种方式明明函数只能调用一次
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1=", res1)
}

func anonDemo2() {
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	res2 := a(10, 30)
	fmt.Println("res2=", res2)
	res3 := a(90, 30)
	fmt.Println("res3=", res3)
}
