package main 

import "fmt"
import "strconv"

func main() {
	num := 123
	snum := strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num, snum)

	num,_ = strconv.Atoi(snum)
	fmt.Println(num)
}