package utils

import "fmt"

var Age int
var Name string

func init() {
	fmt.Println("utils 包的 init()...")
	Age = 100
	Name = "tom..."
}
