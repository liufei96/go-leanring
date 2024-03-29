# 第 11 章面向对象编程(下)

## 11.1 VSCode 的使用

略

## 11.2 面向对象编程思想-抽象

### 11.2.1 抽象的介绍

我们在前面去定义一个结构体时候，实际上就是把一类事物的共有的**属性(字段)**和**行为(方法)**提取出来，形成一个**物理模型**(结构体)。这种研究问题的方法称为抽象。

![image-20231024232420161](./img/image-20231024232420161.png)

### 11.2.2 代码实现

```go
package main

import "fmt"

// 定义一个结构体 Account
type Account struct {
    AccountNo string
    Pwd       string
    Balance   float64
}

// 方法
// 1. 存款
func (account *Account) Deposite(money float64, pwd string) {
    //看下输入的密码是否正确
    if pwd != account.Pwd {
       fmt.Println("你输入的密码不正确!")
       return
    }

    //看看存款金额是否正确
    if money <= 0 {
       fmt.Println("你输入的金额不正确!")
       return
    }

    account.Balance += money
    fmt.Println("存款成功~~")
}

// 取款
func (account *Account) WithDraw(money float64, pwd string) {
    //看下输入的密码是否正确
    if pwd != account.Pwd {
       fmt.Println("你输入的密码不正确!")
       return
    }

    //看看存款金额是否正确
    if money <= 0 || money > account.Balance {
       fmt.Println("你输入的金额不正确!")
       return
    }

    account.Balance -= money
    fmt.Println("取款成功~~")
}

// 查询余额
func (account *Account) Query(pwd string) {
    //看下输入的密码是否正确
    if pwd != account.Pwd {
       fmt.Println("你输入的密码不正确!")
       return
    }
    fmt.Printf("你的账户为=%v, 余额=%v \n", account.AccountNo, account.Balance)
}

func main() {
    //测试一把
    account := Account{
       AccountNo: "gs1111111", Pwd: "666666", Balance: 100.0}
    //这里可以做的更加灵活，就是让用户通过控制台来输入命令... //菜单.... account.Query("666666")
    account.Deposite(200.0, "666666")
    account.Query("666666")
    account.WithDraw(150.0, "666666")
    account.Query("666666")
}
```

对上面代码的要求 

1) 同学们自己可以独立完成 
2) 增加一个控制台的菜单，可以让用户动态的输入命令和选项

### 11.3 面向对象编程三大特性-封装