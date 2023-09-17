package main

import "fmt"

var name = "tom" // 全局变量

//Name := "tom"  // 错误

func test01() {
	fmt.Println(name) // tom
}

func test02() { // 编译器采用就近原则
	name := "jack"
	fmt.Println(name) // jack
}

func main() {
	scopeDemo()

	fmt.Println(name)

	test01() // tom
	test02() // jack
	test01() // tom

}

func scopeDemo() {
	// age 和 Name的作用域只在test函数内部
	age := 10
	Name := "tom..."
	fmt.Println("age=", age, "Name=", Name)
}
