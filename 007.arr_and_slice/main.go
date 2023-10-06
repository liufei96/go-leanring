package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	demo()
	demo2()
	demo3()
	//demo4()
	demo5()
	demo6()

	arr := [3]int{11, 22, 33}
	test01(arr)
	fmt.Println(arr)

	test02(&arr)
	fmt.Println(arr)

	demo7()
	demo8()

	demo9()

	sliceDemo()

	makeDemo()

	makeDemo2()
}

func demo() {
	// 定义六个变量，分别代表六只鸡，然后求出和，然后求出平均值
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0

	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	// 保留小数点2位
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)

	fmt.Printf("totalWeight=%v avgWeight=%v", totalWeight, avgWeight)
}

func demo2() {
	// 定义六个变量，分别代表六只鸡，然后求出和，然后求出平均值
	var hens [7]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	hens[6] = 150.0 // 增加一只鸡

	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	// 保留小数点2位
	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("totalWeight=%v avgWeight=%v", totalWeight, avgWeight)
}

func demo3() {
	var intArr [3]int
	// 当我们定义完数组后，其实数组的各个元素有默认值，默认是0
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p intArr[0]的地址=%p, intArr[1] 地址=%p, intArr[2]地址=%p \n", &intArr, &intArr[0], &intArr[1], &intArr[2])
}

func demo4() {
	// 从终端循环输入5个成绩，保存到float64数组，并输出
	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值\n", i+1)
		fmt.Scanln(&score[i])
	}

	// 变量打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}
}

func demo5() {
	// 四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02=", numArr02)

	// 这里的[...] 是规定的写法
	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04=", numArr04)

	// 类型推导
	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr05=", strArr05)
}

func demo6() {
	heroes := [...]string{"刘备", "关羽", "张飞"}

	for i, value := range heroes {
		fmt.Printf("i=%v v=%v\n", i, value)
		fmt.Printf("heroes[%d]=%v\n", i, heroes[i])
	}

	for _, v := range heroes {
		fmt.Printf("元素的值=%v\n", v)
	}
}

// 创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。使用for 循环访问所有元素并打印出来。提示：字符数据运算 'A'+1 -> 'B'
func demo7() {
	var myChars [26]byte

	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i) // 注意需要将 i => byte
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}
}

func demo8() {
	// 要求：随机生成5个数字，并将其反转打印
	// 思路
	// 1. 随机生成五个数，rand.Intn() 函数
	// 2. 当我们得到随机数后，就放到一个数组 int 数组
	// 3. 反转打印，交换的次数是 len / 2,倒数第一个和第一个元素交换，倒数第2个和第2个元素交换

	var intArr3 [5]int
	// 为了每次生成的随机数不一样，我们需要给一个seed值
	len := len(intArr3)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100) // 0 <= n < 100
	}

	fmt.Println("交换前~=", intArr3)
	// 反转打印，交换的次数是 len / 2
	// 倒数第一个 和 第一个元素交换，倒数第 2 个和第2个元素交换

	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr3[len-1-i]
		intArr3[len-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后~=", intArr3)
}

// 切片的遍历
func demo9() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v \n", i, slice[i])
	}

	fmt.Println()
	// 使用for--range 方式遍历切片
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	// 修改切片后，arr数组的值也会改变
	slice[0] = 21
	fmt.Println(arr)   // [10 21 30 40 50]
	fmt.Println(slice) // [21 30 40]
}

func test01(arr [3]int) {
	arr[0] = 88
}

func test02(arr *[3]int) {
	arr[0] = 88
}

func sliceDemo() {
	// 演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	// 声明/定义一个切片
	// slice := intArr[1:3]
	// 1. slice 就是切片名
	// 2. intArr[1:3] 表示slice 引用到intArr这个数组
	// 3. 引用intArr数组的起始下标为1，最后的下标为3（但是不包括3）
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是=", slice)       // [22 33]
	fmt.Println("slice 的元素个数=", len(slice)) // 2
	fmt.Println("slice 的容量=", cap(slice))   // 切片的容量是可以动态变化的
}

func makeDemo() {
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20

	// 对于切片，必须make使用
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))
}

func makeDemo2() {
	// 方式3
	fmt.Println()

	// 第3种方式：定义一个切片，直接指定具体数组，使用原理类似make的方式
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice size=", len(strSlice)) // 3
	fmt.Println("strSlice cap=", cap(strSlice))  // 3
}
