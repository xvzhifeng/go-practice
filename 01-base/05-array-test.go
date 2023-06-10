package main 

import "fmt"

func main() {
	var a[10]int
	fmt.Println(a)
	for i,_ := range a {
		a[i] = 101
	}
	fmt.Println(a)

	b := [10]float32{1.1,0,1,3}
	fmt.Println(b)
	b[0] = 1
	b[2] = 111
	b[1] = 1222

	fmt.Println(b)

	a1 := [5]int{1,2,3,5,6}
	fmt.Println(a1)
}