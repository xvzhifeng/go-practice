package main 


import "fmt"

func main() {
	var a = 10
	fmt.Printf("%x \n", &a)
	fmt.Println(&a)

	var p *int
	fmt.Println(p)
	p = &a
	fmt.Println(p)
	fmt.Println(*p)
	*p = 20
	fmt.Println(a)
}