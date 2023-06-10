package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["author"] = "sumu"
	m["title"] = "test"
	fmt.Println(m)
	for k,v := range m{
		fmt.Println(k, ": ", v)
	}
	delete(m, "author")
	fmt.Println(m)

	v1, ok := m["author"]
	fmt.Println(v1, ": " , ok)
	v1, ok = m["title"]
	fmt.Println(v1, ": " , ok)
}