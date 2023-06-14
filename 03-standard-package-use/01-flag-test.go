package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	intflag int
	boolflag bool
	stringflag string
)

func init() {
	// 其中函数的四个参数 赋值变量地址 变量名称 默认值 提示语
	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "defualt", "string flag value")
}

func main() {
	flag.Parse()
	
	args := os.Args
	fmt.Println(args)

	fmt.Println("int flag: ", intflag)
	fmt.Println("bool flag: ", boolflag)
	fmt.Println("string flag: ", stringflag)
}