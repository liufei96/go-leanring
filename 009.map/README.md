# 第9章 MAP

## 9.1 map 的基本介绍

map 是 key-value 数据结构，又称为字段或者关联数组。

类似其它编程语言的集合，在编程中是经常使用到

## 9.2 map 的声明

### 9.2.1 基本语法

var map 变量名 map[keytype]valuetype

- **key 可以是什么类型**

  golang 中的 map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 还可以是只包含前面几个类型的 接口, 结构体, 数组

  通常 key 为 int 、string

  **注意: slice， map 还有 function 不可以，因为这几个没法用 == 来判断**

- **valuetype 可以是什么类型**

  valuetype 的类型和 key 基本一样，这里我就不再赘述了

  通常为: 数字(整数,浮点数),string,map,struct

### 9.2.2 map 声明的举例

- map 声明的举例：

  var a map[string]string 

  var a map[string]int 

  var a map[int]string 

  var a map[string]map[string]string 

  **注意：声明是不会分配内存的，初始化需要 make ，分配内存后才能赋值和使用。**

  ```go
  func mapDemo() {
      // map的声明和注意事项
      var a map[string]string
      // 在使用map前，需要先make, make的作用就是给map分配数据空间
      a = make(map[string]string, 10)
      a["no1"] = "刘备"
      a["no2"] = "关羽"
      a["no3"] = "张飞"
      a["no4"] = "诸葛亮"
      fmt.Println(a)
  }
  ```

  对上面代码的说明

  > 1) map 在使用前一定要 make 
  >
  > 2) map 的 key 是不能重复，如果重复了，则以最后这个 key-value 为准
  > 3) map 的 value 是可以相同的. 
  > 4) map 的 key-value 是无序 5) make 内置函数数目

## 9.3 map 的使用

- 第一种方式

  ```go
  func mapDemo() {
  	// map的声明和注意事项
  	var a map[string]string
  	// 在使用map前，需要先make, make的作用就是给map分配数据空间
  	a = make(map[string]string, 10)
  	a["no1"] = "刘备"
  	a["no2"] = "关羽"
  	a["no3"] = "张飞"
  	a["no4"] = "诸葛亮"
  	fmt.Println(a)
  }
  ```

- 第二种方式

  ```go
  func mapDemo2() {
  	// 第二种方式
  	cities := make(map[string]string)
  	cities["no1"] = "杭州"
  	cities["no2"] = "北京"
  	cities["no3"] = "上海"
  	fmt.Println(cities) // map[no1:杭州 no2:北京 no3:上海]
  }
  ```

- 第三种方式

  ```go
  func mapDemo3() {
  	// 第三种方式
  	heroes := map[string]string{
  		"hero1": "李玉芳",
  		"hero2": "典韦",
  		"hero3": "孙悟空",
  	}
  	fmt.Println(heroes) // map[hero1:李玉芳 hero2:典韦 hero3:孙悟空]
  }
  ```

**map 使用的课堂案例**

> 课堂练习：演示一个 key-value 的 value 是 map 的案例 
>
> 比如：我们要存放 3 个学生信息, 
>
> 每个学生有 name 和 sex 信息

```go
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
```

## 9.4 map 的增删改查操作

**map 增加和更新：**

map["key"] = value  //如果 key 还没有，就是增加，如果 key 存在就是修改。

```go
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
```

**map 删除：**

说明： delete(map，"key") ，delete 是一个内置函数，如果 key 存在，就删除该key-value，如果key 不存在，不操作，但是也不会报错

```go
// map的删除
func mapDelete() {
	cities := make(map[string]string)
	cities["no1"] = "杭州"
	cities["no2"] = "北京"
	cities["no3"] = "上海"

	delete(cities, "no1")
	delete(cities, "no4")
	fmt.Println(cities) // map[no2:北京 no3:上海]
}
```

**细节说明**

如果我们要删除 map 的所有 key ,没有一个专门的方法一次删除，可以遍历一下key, 逐个删除或者 map = make(...)，make 一个新的，让原来的成为垃圾，被 gc 回收

```go
// map的删除
func mapDelete() {
    cities := make(map[string]string)
    cities["no1"] = "杭州"
    cities["no2"] = "北京"
    cities["no3"] = "上海"

    // 删除所有的key
    // 1. 遍历所有的key，然后逐一删除
    for key := range cities {
       delete(cities, key)
    }
    // 2. 直接make一个新的孔金
    // cities = make(map[string]string)

    fmt.Println(cities) // map[]
}
```

**map 查找：**

```go 
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
```

**对上面代码的说明:** 

说明：如果 heroes 这个 map 中存在 "no1" ， 那么 findRes 就会返回true,否则返回false

## 9.5 map 遍历：

案例演示相对复杂的 map 遍历：该 map 的 value 又是一个 map 

说明：map 的遍历使用 for-range 的结构遍历

```go
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
```

**map 的长度：**

len(cities)

## 9.6 map 切片

### 9.6.1 基本介绍

**切片的数据类型如果是 map**，则我们称为 slice of map，**map 切片**，这样使用则**map 个数就可以动态变化**了。

### 9.6.2 案例演示

要求：使用一个 map 来记录 monster 的信息 name 和 age, 也就是说一个monster 对应一个map,并且妖怪的个数可以动态的增加=>map 切片
