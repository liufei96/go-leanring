package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 创建一个结构体
	animal := Animal{
		Name:  "黄大仙",
		Age:   "400",
		Score: 87.9,
	}

	//将 Animal 实例传递给 TestStruct 函数
	TestStruct(animal)
}

type Animal struct {
	Name  string `json:"name"`
	Age   string `json:"animal_age"`
	Score float32
	Sex   string
}

// 方法：返回两个数的和
func (s Animal) GetSum(num1 int, num2 int) int {
	return num1 + num2
}

// 方法：接受4个值，给s赋值
func (s Animal) Set(name string, Age string, score float32, sex string) {
	s.Name = name
	s.Age = Age
	s.Score = score
	s.Sex = sex
}

// 方法：显示s的值
func (s Animal) Print() {
	fmt.Println("----- start~ -----")
	fmt.Println(s)
	fmt.Println("----- end~ -----")
}

func TestStruct(a interface{}) {
	// 获取 reflect.Type 类型
	typ := reflect.TypeOf(a)
	// 获取 reflect.Value 类型
	val := reflect.ValueOf(a)
	// 获取a对应的类别，注意这里需要使用val
	kd := val.Kind()
	// 如果传入的不是struct,则退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	// 获取结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	// 变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		// 获取到struct 标签，注意需要通过 reflect.Type 类型来获取
		tagVal := typ.Field(i).Tag.Get("json")
		// 如果该字段于 tag 标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	// 获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	// var params []reflect.Value
	// 方法的排序默认是按照 函数名的排序（ASCII 码）
	val.Method(1).Call(nil) // 获取第二个方法，并调用，这里调用的是print方法

	// 调用结构体的第1个方法 Method(0)
	var params []reflect.Value // 声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}
