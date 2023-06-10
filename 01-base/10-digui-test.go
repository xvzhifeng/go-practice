package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Factorial_no(n uint64) (result uint64) {
	result = 1
	var i uint64 = 1
	for i = 1;i<=n;i++ {
		result *= i
	}
	return
}

func main() {
	fmt.Println(Factorial_no(4))
	fmt.Println(Factorial(4))
}