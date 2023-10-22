package main

import "fmt"

func main() {
	oopDemo1()
	oopDemo2()
	oopDemo3()
	oopDemo4()
	oopDemo5()
	oopDemo6()
	oopDemo7()
	oopDemo8()
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

type Person2 struct {
	Name string
	Age  int
}

func oopDemo4() {
	p2 := Person2{"mary", 20}
	//p2.Name = "tom"
	//p2.Age = 18
	fmt.Println(p2)
}

func oopDemo5() {
	var p3 *Person2 = new(Person2)
	// 因为p3是一个指针，因此标准的给字段赋值的方式
	(*p3).Name = "tom"
	p3.Age = 25 // 这样也行。原因：go的设计者为了程序使用方便，底层会对p3.Age = 25 进行处理。会给 p3 加上取值运算 (*p3).Age = 25
	fmt.Println(p3)
}

func oopDemo6() {
	var person *Person2 = &Person2{}
	// 也可以直接赋值
	// var person *Person2 = &Person2{"tom", 20}
	// 因为p3是一个指针，因此标准的给字段赋值的方式
	//person 是个指针，因此下面访问方式都有可以
	(*person).Name = "lucy"
	person.Age = 20
	fmt.Println(person)
}

type Stu struct {
	Name string
	Age  int
}

func oopDemo7() {
	// 方式1
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明~", 20}

	// 在创建结构体变量时，把名字和字段值写在一起，这种写法，就不依赖字段的定义顺序了
	var stu3 = Stu{
		Name: "jack",
		Age:  19,
	}

	stu4 := Stu{
		Age:  19,
		Name: "lacy",
	}

	fmt.Println(stu1, stu2, stu3, stu4)
}

func oopDemo8() {
	// 方式2,返回结构体指针类型
	var stu5 = &Stu{"小王", 19}
	stu6 := &Stu{"小王~", 20}

	// 在创建结构体变量时，把名字和字段值写在一起，这种写法，就不依赖字段的定义顺序了
	var stu7 = &Stu{
		Name: "小李",
		Age:  19,
	}

	stu8 := &Stu{
		Age:  19,
		Name: "小李～",
	}

	fmt.Println(*stu5, *stu6, *stu7, *stu8)
}
