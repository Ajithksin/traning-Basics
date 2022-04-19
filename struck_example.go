package main

import "fmt"

type person struct {
	first_name string
	age        int
}

func newPerson(name string) *person {
	p := person{first_name: name}
	p.age = 42
	return &p
}
func main() {
	fmt.Println(person{"ajit ", 20})
	s := person{first_name: "abhishek", age: 30}
	fmt.Println(s.first_name)
	sp := &s
	fmt.Println(sp.first_name)
}
