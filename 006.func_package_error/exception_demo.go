package main

import (
	"errors"
	"fmt"
)

func main() {
	exceptionTest()
	fmt.Println("main()下面的代码...")

	readConf2()
}

func exceptionTest() {
	// 使用 defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recober() 内置函数，可以捕获异常
		if err != nil {  // 说明捕获到异常
			// 这里可以将错误信息发送给管理员...
			//fmt.Println("err=", err)
			fmt.Println("发送又叫给admin@sohu.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

func readConf2() {
	err := readConf("config2.ini")
	if err != nil {
		// 如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("readConf2()继续执行")
}
