package main

import "fmt"

func main() {
	ifDemo1()
	ifDemo2()
	ifDemo3()
	ifExp1()
	ifExp2()
	ifExp3()
}

func ifDemo1() {
	//var age int
	//fmt.Println("请输入年龄：")
	//fmt.Scanln(&age)
	//
	//if age > 18 {
	//	fmt.Println("你的年龄大于18，要对自己的行为负责！")
	//}
}

func ifDemo2() {
	if age := 20; age > 18 {
		fmt.Println("你的年龄大于18，要对自己的行为负责！")
	}
}

func ifDemo3() {
	var age = 19
	if age > 18 {
		fmt.Println("你的年龄大于18，要对自己的行为负责！")
	} else {
		fmt.Println("你的年龄不太大这次放过你了")
	}
}

func ifExp1() {
	var x int = 4
	var y int = 1
	if x > 2 {
		if y > 2 {
			fmt.Println(x + y)
		}
		fmt.Println("liufei")
	} else {
		fmt.Println("x is = ", x)
	}
}

func ifExp2() {
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	// 多分支判断
	if score == 100 {
		fmt.Println("奖励一辆BMW")
	} else if score > 80 && score < 99 {
		fmt.Println("奖励一台iphone7plus")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励一台 ipad")
	} else {
		fmt.Println("什么都不奖励")
	}
}

func ifExp3() {
	var second float32
	fmt.Println("请输入秒数：")
	fmt.Scanln(&second)

	if second <= 8 {
		// 进入决赛
		var gender string
		fmt.Println("请输入性别")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入决赛的男子组")
		} else {
			fmt.Println("进入决赛的女子组")
		}
	} else {
		fmt.Println("out...")
	}
}
