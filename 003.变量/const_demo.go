package main

import "fmt"

func main() {

	// 1) 比较简洁的写法
	const (
		a = 1
		b = 1
	)
	fmt.Println(a, b)

	// 2) 还有一种专业的写法
	// iota
	const (
		c = iota
		d
		e
		f
	)
	fmt.Println(c, d, e, f)

	const (
		A    = iota
		B    = iota
		C, D = iota, iota
	)
	fmt.Println(A, B, C, D)

	const (
		A1 = iota
		B2
		_
		C3
	)
	fmt.Println(A1, B2, C3)
}
