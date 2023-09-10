package main

import "fmt"

func main() {
	demo1()
}

func demo1() {
	var num1 int = 10        //声明了 num1
	var num2 int = num1 + 20 //使用 num1
	fmt.Println(num2)
}
