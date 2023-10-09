package main

import "fmt"

func main() {
	oopDemo1()
	oopDemo2()
	oopDemo3()
}

// 张老太养了两只猫猫：一只名字叫小白，今年3岁，白色。还有一只叫小花，今年100岁，花色。
// 请编写一个程序，当用户输入小猫的名字时，显示该猫的名字，年龄，颜色。如果用户输入的小猫名错误，则显示张老太没有这只猫猫。
func oopDemo1() {
	// 1. 使用变量处理
	var cat1Name string = "小白"
	var cat1Age int = 3
	var cat1Color string = "白色"

	var cat2Name string = "小花"
	var cat2Age int = 100
	var cat2Color string = "花色"

	// 2. 使用数组
	var catNames = [...]string{cat1Name, cat2Name}
	var catAges = [...]int{cat1Age, cat2Age}
	var catColors = [...]string{cat1Color, cat2Color}

	fmt.Println(catNames)
	fmt.Println(catAges)
	fmt.Println(catColors)
}

// 定义一个Cat结构体，将Cat的各个字段/属性信息，放入到Cat结构体进行管理
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func oopDemo2() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃<.)))><<"
	fmt.Println("cat1=", cat1)
	fmt.Println("猫猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("Age=", cat1.Age)
	fmt.Println("Color=", cat1.Color)
	fmt.Println("Hobby=", cat1.Hobby)
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              // 指针
	slice  []int             // 切片
	map1   map[string]string // map
}

func oopDemo3() {
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1:", p1.ptr)
	}

	if p1.slice == nil {
		fmt.Println("ok2:", p1.slice)
	}

	if p1.map1 == nil {
		fmt.Println("ok3:", p1.map1)
	}

	// 使用slice前一定要make
	p1.slice = make([]int, 3)
	p1.slice[0] = 100 // ok

	// 使用amp，一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"

	fmt.Println(p1)
}
