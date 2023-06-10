package main

import "fmt"

func main() {
	var (
		name string
		age int
		sex string
	)
	fmt.Scan(&name, &age, &sex)
	fmt.Println(name, age, sex)
	if age <= 18 {
		fmt.Println(name," is small student.")
	} else if 18 < age && age <= 50 {
		fmt.Println(name," is adult.")
	} else {
		fmt.Println(name," is old man.")
	}

	switch sex {
	case "男":
		fmt.Println(name, " 先生 , 你好。")
	case "女":
		fmt.Println(name, " 女士， 你好。")
	default:
		fmt.Println(name, " 你好")
	}
}