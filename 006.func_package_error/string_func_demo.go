package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "hello北"

	// go统一编码时utf-8，utf-8中汉字占用三个字节
	fmt.Println(stringLen(str)) // 8

	stringDemo1(str)

	stringDemo2("122")

	stringDemo3(12345)

	fmt.Println()
	stringDemo4(str)

	stringDemo5([]byte{104, 101, 108, 108, 111, 229, 140, 151})

	stringDemo6()

	stringDemo7()
	fmt.Println()
	stringDemo8()
	fmt.Println()
	stringDemo9()
	fmt.Println()
	stringDemo10()
	stringDemo11()
	stringDemo12()
	stringDemo13()
	stringDemo14()
	stringDemo15()
	stringDemo16()
	stringDemo17()
	stringDemo18()
	stringDemo19()
}

// 统计字符串的长度，按字节 len(str)
func stringLen(str string) int {
	return len(str)
}

// 字符串遍历，同时处理有中文的问题 r := []rune(str)
func stringDemo1(str string) {
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c \n", r[i])
	}
}

// 字符串转整数: n, err := strconv.Atoi("12")
func stringDemo2(str string) {
	n, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("类型转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}
}

// 整数转字符串 str = strconv.Itoa(12345)
func stringDemo3(num int) {
	str := strconv.Itoa(num)
	fmt.Printf("str=%v, str=%T", str, str)
}

// 字符串 转 []byte: var bytes = []byte("hello go")
func stringDemo4(str string) {
	var bytes = []byte(str)
	fmt.Printf("bytes=%v\n", bytes)
}

// []byte 转 字符串: str = string([]byte{97, 98, 99})
func stringDemo5(bytes []byte) {
	str := string(bytes)
	fmt.Println("[]byte转字符串结果:", str)
}

// 10 进制转 2, 8, 16 进制: str = strconv.FormatInt(123, 2) // 2-> 8 , 16
func stringDemo6() {
	var str string
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是=%v\n", str)
}

// 查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
func stringDemo7() {
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b=%v", b)
}

// 统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4
func stringDemo8() {
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v", num)
}

// 不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true
func stringDemo9() {
	b := strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v", b)
}

// 返回子串在字符串第一次出现的 index 值，如果没有返回-1 : strings.Index("NLT_abc", "abc") // 4
func stringDemo10() {
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index=%v \n", index)
}

// 返回子串在字符串最后一次出现的 index，如没有返回-1 : strings.LastIndex("go golang", "go")
func stringDemo11() {
	index := strings.LastIndex("go golang", "go")
	fmt.Printf("LastIndex=%v \n", index) // 3
}

// 将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go 语言", n) n可以指定你希望替换几个，如果 n=-1 表示全部替换
func stringDemo12() {
	str2 := "go go hello"
	str := strings.Replace(str2, "go", "go 语言", -1)
	fmt.Printf("str=%v str2=%v \n", str, str2)
}

// 按 照 指 定 的 某 个 字 符 ， 为 分 割 标 识 ， 将 一 个 字 符 串拆分成字符串数组：strings.Split("hello,wrold,ok", ",")
func stringDemo13() {
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v \n", i, strArr[i])
	}
	fmt.Printf("strArr=%v \n", strArr)
}

// 将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn ")
func stringDemo14() {
	str := strings.TrimSpace(" tn a lone gopher ntrn ")
	fmt.Printf("str=%v \n", str)
}

// 将字符串左右两边指定的字符去掉 ： strings.Trim("! hello! ", " !") // ["hello"] //将左右两边!和 " "去掉
func stringDemo15() {
	str := strings.Trim("! hello! ", " !")
	fmt.Printf("str=%v \n", str)
}

// 将字符串左边指定的字符去掉 ： strings.TrimLeft("! hello! ", " !") // ["hello"] //将左边! 和""去掉
func stringDemo16() {
	str := strings.TrimLeft("! hello! ", " !")
	fmt.Printf("str=%v \n", str)
}

// 将字符串右边指定的字符去掉 ：strings.TrimRight("! hello! ", " !") // ["hello"] //将右边! 和""去掉
func stringDemo17() {
	str := strings.TrimRight("! hello! ", " !")
	fmt.Printf("str=%v \n", str)
}

// 判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
func stringDemo18() {
	flag := strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("flag=%v \n", flag)
}

// 判断字符串是否以指定的字符串结束: strings.HasSuffix("NLT_abc.jpg", "abc") //false
func stringDemo19() {
	flag := strings.HasSuffix("abc.jpg", "abc")
	fmt.Printf("flag=%v \n", flag)
}
