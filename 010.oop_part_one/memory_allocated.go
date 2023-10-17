package main

import "fmt"

type Person3 struct {
	Name string
	Age  int
}

func main() {
	var p1 Person3
	p1.Age = 10
	p1.Name = "小明"

	var p2 *Person3 = &p1 // 这里是关键--> 画出示意图

	//fmt.Println(*p2.Age) // 不能这样写，会报错，因为.的优先级比 * 的高
	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)

	p2.Name = "tom~"
	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name)    // tom~ tom~
	fmt.Printf("p2.Name=%v p1.Name=%v \n", (*p2).Name, p1.Name) // tom~ tom~

	fmt.Printf("p1的地址%p \n", &p1)
	fmt.Printf("p2的地址%p p2的值%p \n", &p2, p2)
}
