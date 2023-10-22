package main

import (
	"encoding/json"
	"fmt"
)

// 结构体
type Point struct {
	x int
	y int
}

// 结构体
type Rect struct {
	leftUp, rightDown Point
}

// 结构体2
type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1 有四个int，在内存中是连续分布的
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// r2 有两个 *Point类型，这个两个 *Point 类型的本身地址也是连续的
	// 但是他们指向的地址不一定连续
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	// 打印地址
	fmt.Printf("r2.leftUp 地址=%p r2.rightDown 地址=%p \n",
		&r2.leftUp, &r2.rightDown)

	// 他们指向的地址不一定是连续的....，这个要看系统运行时时如何分配的
	fmt.Printf("r2.leftUp 指向地址=%p r2.rightDown 指定地址=%p \n",
		r2.leftUp, r2.rightDown)

	test01()
	test02()
	test03()

	var stu Student
	stu.Name = "tom"
	stu.test()

	stu.speak()
	stu.jisaun()
	stu.jisaun2(20)
	res := stu.getSum(10, 20)
	fmt.Println("res=", res)

	var c Circle
	c.radius = 4.0
	res1 := c.area()
	fmt.Println("res1=", res1)

	var i integer = 10
	i.print()
	i.change()
	fmt.Println("main中的i=", i)

	stu1 := Student{
		Name: "zs",
		Age:  20,
	}
	fmt.Println(&stu1)
}

type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v], Age=[%v]", stu.Name, stu.Age)
	return str
}

func (stu Student) test() {
	fmt.Println("test() name=", stu.Name)
}

func (stu Student) speak() {
	fmt.Println(stu.Name, "是一个好人")
}

func (stu Student) jisaun() {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Println(stu.Name, "计算的结果时=", res)
}

func (stu Student) jisaun2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(stu.Name, "计算的结果时=", res)
}

func (stu Student) getSum(n1 int, n2 int) int {
	return n1 + n2
}

//type Stu Student

func test01() {
	var stu1 Student
	var stu2 Stu
	//stu2 = stu1  // 这里是错误的，下面是对的
	stu2 = Stu(stu1)
	fmt.Println(stu1, stu2)
}

type integer int

func test02() {
	var i integer = 10
	var j int = 20
	//j = i // 错误
	j = int(i)
	fmt.Println(i, j)
}

type Monster struct {
	Name  string `json:"name"` // `json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	skill string `json:"skill"`
}

func test03() {
	// 1. 创建一个monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	// 2.将monster变量序列化为json格式字串
	// json.Marshal(monster) 函数中使用反射，这个讲解反射时，会详细介绍
	jsonstr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonstr))
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// Golang 中的方法作用在指定的数据类型上的(即：和指定的数据类型绑定)，因此自定义类型，都可以有方法，而不仅仅是 struct， 比如 int , float32 等都可以有方法

func (i integer) print() {
	fmt.Println("i=", i)
}

// 编写一个方法，可以改变i的值

func (i *integer) change() {
	*i = *i + 1
}
