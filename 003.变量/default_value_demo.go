package main

import "fmt"

func main() {
	defaultValueDemo1()
}

func defaultValueDemo1() {
	var a int          // 0
	var b float32      // 0
	var c float64      // 0
	var isMarried bool // false
	var name string    // ""
	fmt.Printf("a=%d, b=%v, c=%v, isMarried=%v, name=%v", a, b, c, isMarried, name)
}
