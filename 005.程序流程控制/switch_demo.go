package main

import (
	"fmt"
	"strings"
)

func main() {
	//switchDemo1()
	switchDemo2()
	switchDemo4()
	switchDemo3()
	switchDemo5()
	switchDemo6()
	switchDemo7()
	switchDemo8()
	switchDemo9()
	switchDemo10()
}

func switchDemo1() {
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("周一，猴子穿新衣")
	case 'b':
		fmt.Println("周二，猴子当小二")
	case 'c':
		fmt.Println("周三，猴子爬雪山")
	// ...
	default:
		fmt.Println("输入有误")

	}
}

func test(char byte) byte {
	return char
}

func switchDemo2() {
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	switch test(key) {
	case 'a':
		fmt.Println("周一，猴子穿新衣")
	case 'b':
		fmt.Println("周二，猴子当小二")
	case 'c':
		fmt.Println("周三，猴子爬雪山")
	// ...
	default:
		fmt.Println("输入有误")

	}
}

func switchDemo3() {
	/*	var n1 int32 = 20
		var n2 int64 = 20
		switch n1 {
		case n2: //错误，原因是 n2的数据类型和n1不一致
			fmt.Println("ok1")
		}*/
}

// case 后面可以有多个表达式
func switchDemo4() {
	var n1 int32 = 5
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5: // case 后面可以有多个表达式
		fmt.Println("ok1")
	default:
		fmt.Println("没有匹配到")
	}
}

// case 后面的表达式如果是常量值(字面量)，则要求不能重复
func switchDemo5() {
	var n1 int32 = 5
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5: // case 后面可以有多个表达式
		fmt.Println("ok1")
	//case 5: // 错误，因为前面我们有常量5，因此重复，就会报错
	//	fmt.Println("ok2~")
	default:
		fmt.Println("没有匹配到")
	}
}

// switch 后面也可以不带表达式
func switchDemo6() {
	var age = 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到")
	}
}

// switch 后也可以直接声明/定义一个变量，分号结束，不推荐
func switchDemo7() {
	switch grade := 90; {
	case grade > 90:
		fmt.Println("成绩优秀~..")
	case grade >= 70 && grade <= 90:
		fmt.Println("成绩优良~..")
	case grade >= 60 && grade < 70:
		fmt.Println("成绩优良~..")
	default:
		fmt.Println("不及格~")
	}
}

// switch 穿透-fallthrough ，如果在 case 语句块后增加 fallthrough ,则会继续执行下一个case，也叫 switch 穿透
func switchDemo8() {
	var num = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // 默认只能穿透一层
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到..")
	}
}

var x interface{}

func switchDemo9() {
	var y = 10.0
	x = y

	switch i := x.(type) {
	case nil:
		fmt.Println("x 的类型 ~: %T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}
}

func switchDemo10() {
	var char byte
	fmt.Println("请输入一个字符...")
	fmt.Scanf("%c", &char)

	switch char {
	case 'a':
		fmt.Println(strings.ToUpper("a"))
	case 'b':
		fmt.Println(strings.ToUpper("b"))
	case 'c':
		fmt.Println(strings.ToUpper("c"))
	case 'd':
		fmt.Println(strings.ToUpper("d"))
	case 'e':
		fmt.Println(strings.ToUpper("e"))
	default:
		fmt.Println("other")
	}
}
