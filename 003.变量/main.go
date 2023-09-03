package main

import "fmt"

// 定义全局变量
var n1 = 100
var n2 = 200
var name = "liufei"

// 上面的声明方式，也可以改成一次性声明
var (
	n3    = 300
	n4    = 900
	name2 = "mary"
)

func main() {
	getVal(1, 3)
	test1()
	test1Var1()
	test1Var2()
	test1Var3()
	moreVar()
	changeVar()
	defaultVar()
}

func test1() {
	// 定义变量
	var i int
	// 给变量赋值
	i = 10
	// 使用变量
	fmt.Println(i)
}

func test1Var1() {
	// 第一种：声明后不赋值，使用默认值
	var i int
	// 使用变量
	fmt.Println("i=", i)
}

func test1Var2() {
	// 第二种：根据值自行判断类型
	var num = 10.11
	// 使用变量
	fmt.Println("num=", num)
}

func test1Var3() {
	// 第三种：省略 var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	name := "liufei"
	// 使用变量
	fmt.Println("name=", name)
}

func moreVar() {
	// 声明多个变量
	//var n1, n2, n3 int
	//fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 一次性声明多个变量的方式2
	//var n1, name, n3 = 100, "tom", 888
	//fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	// 一次性声明多个变量的方式3
	n1, name, n3 := 100, "tom=", 888
	fmt.Println("n1=", n1, "name=", name, "n3=", n3)
}

func changeVar() {
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=", i)
	// i = 1.2 // int ，议案因不能改变数据类型
}

// 测试默认值
func defaultVar() {
	var i int
	var bol bool
	var name string
	var flo = float32(0.0)
	fmt.Println("i=", i)       // 0
	fmt.Println("bol=", bol)   // false
	fmt.Println("name=", name) // 空串
	fmt.Println("flo=", flo)   // 0
}

// 一个函数可以返回多个值
func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}
