package main

import "fmt"

func main() {
	var price float32 = 89.12
	fmt.Println("price=", price)

	floatDemo1()
	floatDemo2()
	floatDemo3()
	floatDemo4()
}

func floatDemo1() {
	var num1 float32 = -0.0089
	var num2 float64 = -7809656.09
	fmt.Println("num1=", num1, "num2=", num2)
}

// 造成精度损失
func floatDemo2() {
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4) // num3= -123.00009 num4= -123.0000901
}

// 默认是float64
func floatDemo3() {
	var num5 = -123.0000901
	fmt.Printf("num5的数据类型是 %T \n", num5) // num5的数据类型是 float64
}

// 浮点型常量的表示形式
func floatDemo4() {
	num6 := 5.12
	num7 := .123                              // => 0.123
	fmt.Println("num6=", num6, "num7=", num7) // num6= 5.12 num7= 0.123

	// 科学计数法
	num8 := 5.1234e2                                           // ? 5.1234 * 10 的2次方
	num9 := 5.1234e2                                           // ? 5.1234 * 10 的2次方
	num10 := 5.1234e-2                                         // ? 5.1234 / 10 的2次方
	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10) // num8= 512.34 num9= 512.34 num10= 0.051234
}
