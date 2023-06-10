package main

import "fmt"
import "time"

func print(n int, s string, c chan int) {
	sum := 0
	for i:=1 ; i<n;i++ {
		sum += i
		fmt.Println(i, s)
		time.Sleep(100*time.Millisecond)
	}
	c<-sum
}

func main () {
	c := make(chan int)
	go print(5, "t01", c)
	go print(6, "t02", c)
	time.Sleep(1000*time.Millisecond)
	a := <- c
	b := <- c
	fmt.Println(a, b)
	fmt.Println("done")
}