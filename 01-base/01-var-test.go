package main

import "fmt"
// Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。

// 声明变量的一般形式是使用 var 关键字：

// var identifier type

func main() {
	// 如果定义的变量为赋值，则int型默认值为0
	var age int
	fmt.Println(age)
	var height int = 175
	fmt.Println(height)
	// 如果未指定类型，则会更具初始值推到出类型
	var weight = "120"
	fmt.Println(weight)
	fmt.Printf("%t \n",weight)

	// 常用
	name := "test01"
	fmt.Println(name)

	// var 代表变量 const 代表常量

	const arm = 2
	fmt.Println(arm)

	var x, y int
	var (  // 这种因式分解关键字的写法一般用于声明全局变量
		a int
		b bool
	)

	var c, d int = 1, 2
	var e, f = 123, "hello"

	//这种不带声明格式的只能在函数体中出现
	//g, h := 123, "hello"
	g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h)
}