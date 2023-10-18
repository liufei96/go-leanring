package main

import "fmt"

func main() {
	var mu MethodUtils
	mu.Print()
	mu.PrintMulTable(8)

	var arr = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	ans := mu.transpose(arr)
	fmt.Println(ans)
}

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) PrintMulTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func (mu MethodUtils) transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	var res = make([][]int, cols)
	for i := 0; i < cols; i++ {
		res[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}
