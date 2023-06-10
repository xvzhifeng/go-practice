package main

import "fmt"

type Animal interface {
	voice()
	action()
}

type Duck struct {

}

type Dog struct {

}

type Pig struct {

}

func (duck Duck) voice() {
	fmt.Println("ga ga ga ...")
}
func (dog Dog) voice() {
	fmt.Println("wang wang wang ...")
}
func (pig Pig) voice() {
	fmt.Println("no no no ...")
}

func (duck Duck) action() {
	fmt.Println("da da da ...")
}
func (dog Dog) action() {
	fmt.Println("bark bark ...")
}
func (pig Pig) action() {
	fmt.Println("eat eat ...")
}




func main() {
	var animal Animal
	animal = new(Duck)
	animal.voice()

	animal = new(Pig)
	animal.voice()

}