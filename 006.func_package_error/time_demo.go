package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T \n", now, now)

	// 2. 通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 3. 格式化日期
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dataStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dataStr=%v\n", dataStr)

	// 3.1 使用date.Format()
	fmt.Printf(now.Format("2006-01-02 15:04:05")) // 这里面的每个数字都是固定的
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	//timeDemo()

	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())

	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03()耗费时间为%v秒\n", end-start)
}

func timeDemo() {
	// 需求：每隔1秒打印出一个数字，打印到100就退出
	// 需求2：每隔0.1秒打印出一个数字，打印到100就退出
	for i := 1; i <= 100; i++ {
		// time.Sleep(1 * time.Second) // 1 秒
		time.Sleep(100 * time.Millisecond) // 0.1 秒
		fmt.Println(i)
	}
}

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
