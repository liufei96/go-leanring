package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	m := Monster{
		Name: "玉兔精",
		Age:  20,
		Sal:  888.99,
		Sex:  "female",
	}
	data, _ := json.Marshal(m)
	fmt.Println("json result:", string(data))

	// 请编写一个案例，演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作
	var num int = 100
	reflectTest01(num)

	//2. 定义一个 Student 的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}

type Monster struct {
	Name string  `json:"monsterName"`
	Age  int     `json:"monsterAge"`
	Sal  float64 `json:"monsterSal"`
	Sex  string  `json:"monsterSex"`
}

// 专门演示反射
func reflectTest01(b interface{}) {
	// 通过反射获取的传入的变量的 type , kind, 值

	// 1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	// 2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	// 这里不能直接 2 + rVal，rVal是reflect.Value类型
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	// 下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()

	// 将interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

// 专门演示反射[对结构体的反射]
func reflectTest02(b interface{}) {
	// 通过反射获取的传入的变量的 type , kind, 值
	// 1. 先获取到 reflect.Type
	vTyp := reflect.TypeOf(b)
	fmt.Println("vTyp=", vTyp)

	// 2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)

	// 3. 获取 变量对应的kind
	// (1) rVal.Kind() =>
	kind1 := rVal.Kind()
	// (2) rType.Kind() =>
	kind2 := vTyp.Kind()
	fmt.Printf("kind1=%v, kind2=%v \n", kind1, kind2) // 两个结果都是一样，都是struct

	// 下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)

	// 将 interface{} 通过断言转成需要的类型
	// 这里，我们就简单使用了一带检测的类型断言
	// 同学们可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}
