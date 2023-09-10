# 第六章：函数、包和错误处理

## 1. 为什么需要函数

### 1.1 请大家完成这样一个需求:

输入两个数,再输入一个运算符(+,-,*,/)，得到结果.。

### 1.2 使用传统的方法解决

走代码

```go
func main() {
    var n1 float64 = 1.2
    var n2 float64 = 2.3
    var operator byte = '-'
    var res float64
    switch operator {
    case '+':
       res = n1 + n2
    case '-':
       res = n1 - n2
    case '*':
       res = n1 * n2
    case '/':
       res = n1 / n2
    default:
       fmt.Printf("操作符号错误...")
    }
    fmt.Println("res=", res)
}
```

分析一下上面代码问题

1) 上面的写法是可以完成功能, 但是代码冗余 
2) 同时不利于代码维护 
3) 函数可以解决这个问题

## 2. 函数的基本概念

为完成某一功能的程序指令(语句)的集合,称为函数。 

在 Go 中,函数分为: 自定义函数、系统函数(查看 Go 编程手册)

## 3. 函数的基本语法

```go
func 函数名(形参列表)(返回值列表) {
	执行语句...
	return 返回值列表
}
```

1 形参列表：表示函数的输入

2 函数中的语句：表示为了实现某一功能代码块

3 函数可以有返回值，也可以没有

## 4. 快速入门案例

使用函数解决前面的计算问题。 

走代码:

```go
func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Printf("操作符号错误...")
	}
	return res
}
```

调用函数

```go
func main() {
    var n1 float64 = 1.2
    var n2 float64 = 2.3
    var operator byte = '-'
    fmt.Println("res=", cal(n1, n2, operator))
}
```

## 5. 包的引出

在实际的开发中，我们往往需要在不同的文件中，去调用其它文件的定义的函数，比如main.go中，去使用 utils.go 文件中的函数，如何实现？ -》包

现在有两个程序员共同开发一个 Go 项目，程序员 xiaoming 希望定义函数Cal ，程序员xiaoqiang也想定义函数也叫 Cal。两个程序员为此还吵了起来,怎么办? -》包

## 6. 包的原理图

包的本质实际上就是创建不同的文件夹，来存放程序文件。

## 7. 包的基本概念

说明：go 的每一个文件都是属于一个包的，也就是说 go 是以包的形式来管理文件和项目目录结构的

## 8. 包的三大作用

区分相同名字的函数、变量等标识符

当程序文件很多时,可以很好的管理项目 

控制函数、变量等访问范围，即作用域

## 9. 包的相关说明

### 打包基本语法

```go
package 包名
```

### 引入包的基本语法

```go
import "包的路径"
```

## 10. 包使用的快速入门

包快速入门-Go 相互调用函数，我们将 func Cal 定义到文件 utils.go , 将utils.go 放到一个包中，当其它文件需要使用到 utils.go 的方法时，可以 import 该包，就可以使用了. 

【为演示：新建项目目录结构】

![image-20230910231732913](./img/包项目演示.png)

utils.go 文件

```go
package utils

import "fmt"

func Cal(n1 float64, n2 float64, operator byte) float64 {
    var res float64
    switch operator {
    case '+':
       res = n1 + n2
    case '-':
       res = n1 - n2
    case '*':
       res = n1 * n2
    case '/':
       res = n1 / n2
    default:
       fmt.Printf("操作符号错误...")
    }
    return res
}
```

func_demo.go  文件

```go
package main

import (
    "fmt"
    "go-leanring/006.func_package_error/utils"
)

func main() {
    var n1 float64 = 1.2
    var n2 float64 = 2.3
    var operator byte = '-'
    fmt.Println("res=", utils.Cal(n1, n2, operator))
}
```

注意：

1. 整个go-leanring 工程需要放到GO安装目录下的src目录下，不然包引用无效
2. 包名称不能包含中文，所以这一章目录使用英文名

## 11. 包使用的注意事项和细节讨论

1) 在给一个文件打包时，该包对应一个文件夹，比如这里的 utils 文件夹对应的包名就是utils，文件的包名通常和文件所在的文件夹名一致，一般为小写字母。 

2) 当一个文件要使用其它包函数或变量时，需要先引入对应的包

   引入方式 1：import "包名"

   引入方式 2：

   ```go
   import ( "包名"
   	"包名"
   )
   ```

   package 指令在 文件第一行，然后是 import 指令。

   在 import 包时，路径从 $GOPATH 的 src 下开始，不用带 src , 编译器会自动从src 下开始引入

3. 为了让其它包的文件，可以访问到本包的函数，则该函数名的**首字母**需要大写，类似其它语言的 public ,这样才能跨包访问。比如 `utils.go` 的Cal函数
4. 在访问其它包函数，变量时，其语法是 包名.函数名， 比如这里的 `func_demo.go`文件中

5. 如果包名较长，Go 支持给包取别名， 注意细节：取别名后，原来的包名就不能使用了

   ```go
   import (
   	"fmt"
   	util "go-leanring/006.func_package_error/utils"
   )
   
   func main() {
   	var n1 float64 = 1.2
   	var n2 float64 = 2.3
   	var operator byte = '-'
       // 这里使用别名 util 调用
   	fmt.Println("res=", util.Cal(n1, n2, operator))
   }
   ```

   说明: 如果给包取了别名，则需要使用别名来访问该包的函数和变量

6) 在同一包下，不能有相同的函数名（也不能有相同的全局变量名），否则报重复定义
7) 如果你要编译成一个可执行程序文件，就需要将这个包声明为 main , 即package main.这个就是一个语法规范，如果你是写一个库 ，包名可以自定义

## 12. 函数的调用机制

### 12.1 通俗易懂的方式的理解

![image-20230910232541826](./img/函数的调用机制.png)

### 12.2 函数-调用过程

**传入一个数+1**

![image-20230910232738480](./img/image-20230910232738480.png)

对上图说明

(1) 在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间和其它的栈的空间区分开来

(2) 在每个函数对应的栈中，数据空间是独立的，不会混淆

(3) 当一个函数调用完毕(执行完毕)后，程序会销毁这个函数对应的栈空间。

**计算两个数,并返回**

```go
package main

import (
	"fmt"
	util "go-leanring/006.func_package_error/utils"
)

func main() {
    var n1 float64 = 1.2
    var n2 float64 = 2.3
    var operator byte = '-'
    fmt.Println("res=", util.Cal(n1, n2, operator))

    n3 := 10
    test(n3)

    fmt.Println("main() n1=", n1) // ? 输出结果=?

    sum := getSum(10, 20)
    fmt.Println("main sum = ", sum)
}

func test(n1 int) {
    n1 = n1 + 1
    fmt.Println("test() n1 = ", n1) // 输出结果
}

func getSum(n1 int, n2 int) int {
    sum := n1 + n2
    fmt.Println("getSum() sum = ", sum)
    return sum
}
```

## 12.3 return 语句