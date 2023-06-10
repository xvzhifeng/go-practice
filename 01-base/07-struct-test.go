package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var b1 = Books{"test","sumu","computer",0001}
	fmt.Println(b1)
	b1.title = "test 02"
	fmt.Println(b1)

	var bp *Books
	bp = &b1
	fmt.Println(bp)
	fmt.Println(&b1)
	bp.book_id = 00002
	fmt.Println(b1)
}