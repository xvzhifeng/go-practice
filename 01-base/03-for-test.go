package main

import "fmt"

func main() {
	for i := 2; i < 100; i++ {
		var flag = false
		for j := 2; j < i ;j++ {
			if i % j == 0 {
				flag = true
				break
			}
		}
		if !flag {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	zfb99()
	zfb99Reverse()

}


func zfb99() {
	for i := 1; i<=9; i++ {
		for j := 1 ; j<=i ; j++ {
			fmt.Print(j, " x ", i , " = ", j*i, "  ")
		}
		fmt.Println()
	}
}


func zfb99Reverse() {
	for i := 1; i<=9; i++ {
		for j := 9 ; j>=i ; j-- {
			fmt.Print(j, " x ", i , " = ", j*i, "  ")
		}
		fmt.Println()
	}
}

