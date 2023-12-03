package main

import (
	"reflect"
	"testing"
)

// 案例3: 定义了两个函数 test1 和 test2，定义一个适配器函数用作统一处理接口【了解】
func TestReflectFunc(t *testing.T) {
	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)

	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}

	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}

	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}

	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}

type user struct {
	UserId string
	Name   string
}

// 4，使用反射操作任意结构体类型：【了解】
func TestReflectStruct(t *testing.T) {
	var (
		model *user
		sv    reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String())
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem", sv.Kind().String())
	sv.FieldByName("UserId").SetString("123")
	sv.FieldByName("Name").SetString("zhangsan")
	t.Log("model", model)
}

// 5. 使用反射创建并操作结构体
func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                       // 获取类型 *user
	t.Log("reflect.ValueOf", st.Kind().String())     // ptr
	st = st.Elem()                                   // st指向的类型
	t.Log("reflect.TypeOf.Elem", st.Kind().String()) // struct
	// New返回一个Value类型值，该值持有一个指向类型为type的新申请的零值的指针
	elem = reflect.New(st)
	t.Log("reflect.New", elem.Kind().String())
	t.Log("reflect.New.Elem", elem.Elem().Kind().String())
	// model就是创建的user结构体变量（实例）
	model = elem.Interface().(*user) // model是*user 它的指向和elem是一样的
	elem = elem.Elem()               // 取得elem指向的值
	elem.FieldByName("UserId").SetString("456")
	elem.FieldByName("Name").SetString("lisi")
	t.Log("model model.Name", model, model.Name)
}
