package main

import "fmt"

func main() {
	forDemo1()
	forDemo2()
	forDemo3()
	forDemo4()
	forDemo5()
	forDemo6()
	forDemo7()
}

func forDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Println("你好，GO语言！")
	}
}

// for 循环的第二种看法
func forDemo2() {
	j := 1
	for j <= 10 {
		fmt.Println("你好，GO语言-2")
		j++
	}
}

// for 循环的第三种看法
func forDemo3() {
	k := 1
	for {
		if k <= 10 {
			fmt.Println("ok!!")
		} else {
			break
		}
		k++
	}
}

// 遍历字符串-传统方法
func forDemo4() {
	var str string = "hello,world!" // 这里如果有中文会乱码
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}
}

// 遍历字符串-for-range
func forDemo5() {
	fmt.Println()
	str := "abcdefg杭州"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c  \n", index, val)
	}
}

// 遍历字符串 切片
func forDemo6() {
	fmt.Println()
	str := "abcdefg杭州"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}
}

// 案例: 打印 1~100 之间所有是 9 的倍数的整数的个数及总和
func forDemo7() {
	var count = 0
	var sum = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v, sum=%v", count, sum)
}
