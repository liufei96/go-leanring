package main

import "fmt"

func main() {
	stringDemo1()
	stringDemo2()

	fbnSlice := fbn(10)
	fmt.Println("fbnSlice", fbnSlice)
}

// string 底层是一个 byte 数组，因此 string 也可以进行切片处理
func stringDemo1() {
	str := "hello杭州"
	slice := str[2:]
	fmt.Println("slice=", slice) // llo杭州

	// str[0] = 'z' // 编译不会通过
}

// 如果需要修改字符串，可以先将 string -> []byte / 或者 []rune -> 修改-> 重写转成string
func stringDemo2() {
	str := "hello杭州"
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	// 细节，我们转换成[]byte后，可以处理英文和数字，但是不能处理中文
	// 原因是 []byte 字节来处理，而一个汉字，是3个字节，因此就会出现乱码
	// 解决方法是将 string 转成 []rune 即可，因为 []rune 是按字符处理，兼容汉字
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr1)
	fmt.Println("str=", str)
}

func fbn(n int) []uint64 {
	// 声明一个切片
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
