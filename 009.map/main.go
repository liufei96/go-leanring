package main

import "fmt"

func main() {
	mapDemo()
	mapDemo2()
	mapDemo3()
	mapDemo4()
	mapAddAndUpdate()
	mapDelete()
	makeQuery()
	mapDemo5()
}

func mapDemo() {
	// map的声明和注意事项
	var a map[string]string
	// 在使用map前，需要先make, make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "刘备"
	a["no2"] = "关羽"
	a["no3"] = "张飞"
	a["no4"] = "诸葛亮"
	fmt.Println(a) // map[no1:刘备 no2:关羽 no3:张飞 no4:诸葛亮]
}

func mapDemo2() {
	// 第二种方式
	cities := make(map[string]string)
	cities["no1"] = "杭州"
	cities["no2"] = "北京"
	cities["no3"] = "上海"
	fmt.Println(cities) // map[no1:杭州 no2:北京 no3:上海]
}

func mapDemo3() {
	// 第三种方式
	heroes := map[string]string{
		"hero1": "李玉芳",
		"hero2": "典韦",
		"hero3": "孙悟空",
	}
	fmt.Println(heroes) // map[hero1:李玉芳 hero2:典韦 hero3:孙悟空]
}

// 我们要存放 3 个学生信息,
// 每个学生有 name 和 sex 信息
func mapDemo4() {
	studentMp := make(map[string]map[string]string)

	studentMp["stu01"] = make(map[string]string, 3)
	studentMp["stu01"]["name"] = "tom"
	studentMp["stu01"]["sex"] = "男"
	studentMp["stu01"]["address"] = "杭州"

	studentMp["stu02"] = make(map[string]string, 3) // 这句不能少
	studentMp["stu02"]["name"] = "mary"
	studentMp["stu02"]["sex"] = "女"
	studentMp["stu02"]["address"] = "上海黄浦江..."

	fmt.Println(studentMp)                     // map[stu01:map[address:杭州 name:tom sex:男] stu02:map[address:上海黄浦江... name:mary sex:女]]
	fmt.Println(studentMp["stu01"])            // map[address:杭州 name:tom sex:男]
	fmt.Println(studentMp["stu02"]["address"]) // 上海黄浦江...
}

// map的增肌和更新
func mapAddAndUpdate() {
	cities := make(map[string]string)
	cities["no1"] = "杭州"
	cities["no2"] = "北京"
	cities["no3"] = "上海"

	// no1 已经存在，则更新
	cities["no1"] = "深圳"

	fmt.Println(cities) // map[no1:深圳 no2:北京 no3:上海]
}

// map的删除
func mapDelete() {
	cities := make(map[string]string)
	cities["no1"] = "杭州"
	cities["no2"] = "北京"
	cities["no3"] = "上海"

	delete(cities, "no1")
	delete(cities, "no4")
	fmt.Println(cities) // map[no2:北京 no3:上海]

	// 删除所有的key
	// 1. 遍历所有的key，然后逐一删除
	for key := range cities {
		delete(cities, key)
	}
	// 2. 直接make一个新的孔金
	// cities = make(map[string]string)

	fmt.Println(cities) // map[]
}

// make的查找
func makeQuery() {
	cities := make(map[string]string)
	cities["no1"] = "杭州"
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有no1 key 的值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key \n")
	}
}

// 使用 for range 遍历 map
func mapDemo5() {
	cities := make(map[string]string)
	cities["no1"] = "杭州"
	cities["no2"] = "北京"
	cities["no3"] = "上海"

	for k, v := range cities {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
}
