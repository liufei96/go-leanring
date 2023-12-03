package main

import (
	"fmt"
	"reflect"
)

func main() {
	cal := Cal{
		Num1: 8,
		Num2: 3,
	}
	testReflectStruct(cal)
}

type Cal struct {
	Num1 int32 `json:"num1"`
	Num2 int32 `json:"num2"`
}

func (cl Cal) GetSub(name string) {
	res := cl.Num1 - cl.Num2
	fmt.Printf("%s 完成了减法运行，%d - %d = %d ", name, cl.Num1, cl.Num2, res)
}

func testReflectStruct(b interface{}) {
	// 获取Cal的字段信息
	typ := reflect.TypeOf(b)

	val := reflect.ValueOf(b)
	fieldNum := val.NumField()
	fmt.Printf("struct has %d fields\n", fieldNum)
	for i := 0; i < fieldNum; i++ {
		// 获取到struct 标签，注意需要通过 reflect.Type 类型来获取
		tagVal := typ.Field(i).Tag.Get("json")
		// 如果该字段于 tag 标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))
	val.MethodByName("GetSub").Call(params)
}
