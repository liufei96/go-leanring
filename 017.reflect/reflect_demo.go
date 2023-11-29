package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 10
	testInt(&num)
	fmt.Println("num=", num)

	var n int = 100
	fn := reflect.ValueOf(&n)
	fn.Elem().SetInt(200)
	fmt.Printf("n=%v\n", n) // n=200

	// fn.Elem() 用于获取指针指向变量，类似下面
	var num1 = 100
	var b *int = &num1
	*b = 3
	fmt.Println("num1=", num1) // num1= 3

	fmt.Println("")

	var v float64 = 1.2
	expDemo1(v)

	expDemo2()
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val type=%T\n", val)
	val.Elem().SetInt(20)
	fmt.Printf("val=%v\n", val) // 这里val返回的是地址
}

func expDemo1(b interface{}) {
	// 获取Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp) // rType= float64

	rKind := rTyp.Kind()
	fmt.Println("rKind=", rKind) // rKind= float64

	// 获取rVal
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal) // rVal= 1.2

	rKind2 := rVal.Kind()
	fmt.Println("rKind2=", rKind2) // rKind2= float64

	// 将reflect.Value转成interface
	iV := rVal.Interface()
	fmt.Println("iV=", iV) // iV= 1.2

	// 再将interface{} 转成float
	val := iV.(float64)
	fmt.Println("val=", val) // val= 1.2
}

func expDemo2() {
	var str string = "tom"
	//fs := reflect.ValueOf(str)
	//fs.SetString("jack")

	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")

	fmt.Printf("%v\n", str)
}
